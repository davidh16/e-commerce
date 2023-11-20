package server

import (
	"e-commerce/models"
	"encoding/json"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"strings"
	"time"
)

func (s *Server) Register(w http.ResponseWriter, req *http.Request) {

	var user *models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	user.RoleUuid = "9c8e2b1e-b839-4e49-937e-d460065eccb6"

	// calling service to create user
	_, token, err := s.service.CreateUser(user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	err = s.service.SendVerificationEmail(user.EmailAddress, user.FirstName, token.Token)
	if err != nil {
		log.Print("verification email have not been sent")
		return
	}

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

func (s *Server) VerifyAccount(w http.ResponseWriter, req *http.Request) {
	var tokenString string
	err := json.NewDecoder(req.Body).Decode(&tokenString)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	token, err := s.service.GetToken(tokenString)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	if token.IsUsed != false {
		returnResponse(w, http.StatusBadRequest, errors.New("tokes had been already used"), nil)
		return
	}

	if time.Now().After(token.ExpiresAt) {
		returnResponse(w, http.StatusBadRequest, errors.New("tokes had expired"), nil)
		return
	}

	_, err = s.service.UpdateVerificationToken(token)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	user, err := s.service.GetUser(token.UserUuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	_, err = s.service.VerifyUser(user)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return

}
