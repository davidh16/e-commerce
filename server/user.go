package server

import (
	"e-commerce/models"
	"encoding/json"
	"errors"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func (s *Server) Register(w http.ResponseWriter, req *http.Request) {

	var user *models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(req.Context(), w, http.StatusBadRequest, err)
		return
	}

	// calling service to create user
	_, err = s.service.CreateUser(req.Context(), user)
	if err != nil {
		returnResponse(req.Context(), w, http.StatusBadRequest, err)
		return
	}

	//TODO implement sending verification email

	returnResponse(req.Context(), w, http.StatusCreated, nil)
	return
}

func (s *Server) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(req.Context(), w, http.StatusBadRequest, err)
		return
	}

	// validating given credentials (if ok, user model is retrieved)
	dbUser, err := s.service.ValidateCredentials(user)
	if err != nil {
		if err.Error() == "record not found" {
			returnResponse(req.Context(), w, http.StatusNotFound, err)
			return
		}
		returnResponse(req.Context(), w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	var accessToken string
	var refreshToken string

	g := new(errgroup.Group)
	g.Go(func() error {
		// generating access token for authenticated user
		accessToken, err = s.service.GenerateJWT(dbUser.Uuid, false)
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
		refreshToken, err = s.service.GenerateJWT(dbUser.Uuid, true)
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
		returnResponse(req.Context(), w, http.StatusInternalServerError, err)
		return
	}

	// creating  json response
	response, err := json.Marshal(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
	if err != nil {
		returnResponse(req.Context(), w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
