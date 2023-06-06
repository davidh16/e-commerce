package controller

import (
	"e-commerce/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (c Controller) Register(w http.ResponseWriter, req *http.Request) {

	var user models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err)
		return
	}

	// calling service to create user
	_, err = c.service.Create(user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err)
		return
	}

	//TODO implement sending verification email

	returnResponse(w, http.StatusCreated, nil)
	return
}

func (c Controller) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err)
		return
	}

	// validating given credentials (if ok, user model is retrieved)
	dbUser, err := c.service.ValidateCredentials(user)
	if err != nil {
		if err.Error() == "record not found" {
			returnResponse(w, http.StatusNotFound, err)
			return
		}
		returnResponse(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	// generating token for authorized user
	token, err := c.service.GenerateJWT(dbUser.Uuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err)
		return
	}

	// writing a token in body of response message
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(token))
	return
}
