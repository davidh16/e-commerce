package server

import (
	"e-commerce/models"
	"encoding/json"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
)

func (s *Server) Register(w http.ResponseWriter, req *http.Request) {

	var user *models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	// calling service to create user
	_, err = s.service.CreateUser(user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	//TODO implement sending verification email

	returnResponse(w, http.StatusCreated, nil, nil)
	return
}

func (s *Server) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	// validating given credentials (if ok, user model is retrieved)
	dbUser, err := s.service.ValidateCredentials(user)
	if err != nil {
		if err.Error() == "record not found" {
			returnResponse(w, http.StatusNotFound, err, nil)
			return
		}
		returnResponse(w, http.StatusUnauthorized, errors.New("invalid credentials"), nil)
		return
	}

	var accessToken string
	var refreshToken string

	g := new(errgroup.Group)
	g.Go(func() error {
		// generating access token for authenticated user
		accessToken, err = s.service.GenerateJWT(dbUser, false)
		if err != nil {
			return err
		}

		// saving access token in memory (redis)
		err = s.service.SaveAccessToken(dbUser.Uuid, accessToken)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		// generating refresh token for authenticated user
		refreshToken, err = s.service.GenerateJWT(dbUser, true)
		if err != nil {
			return err
		}

		// saving refresh token in database
		err = s.service.SaveRefreshToken(refreshToken)
		if err != nil {
			return err
		}

		return nil
	})

	if err = g.Wait(); err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	// creating  json response
	var response struct {
		AccessToken  string
		RefreshToken string
	}
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	returnResponse(w, http.StatusOK, err, response)
	return
}

func (s *Server) Me(w http.ResponseWriter, req *http.Request) {
	reqToken := req.Header.Get("Authorization")
	me, err := s.service.Me(reqToken)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, me)
	return
}

func (s *Server) Logout(w http.ResponseWriter, req *http.Request) {
	reqToken := req.Header.Get("Authorization")
	reqToken = strings.Split(reqToken, "Bearer ")[1]

	err := s.service.BlackListToken(reqToken)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, nil)
	return
}
