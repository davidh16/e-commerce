package controller

import (
	"e-commerce/models"
	"encoding/json"
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

	returnResponse(w, http.StatusCreated, nil)
	return
}
