package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) SaveAddress(w http.ResponseWriter, req *http.Request) {
	var address *models.Address

	err := json.NewDecoder(req.Body).Decode(&address)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.CreateAddress(address)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, nil)
	return
}

func (s *Server) GetAddress(w http.ResponseWriter, req *http.Request) {
	var request struct {
		AddressUuid string `json:"address_uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	address, err := s.service.FindAddressesByUuid(request.AddressUuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, address)
	return
}

func (s *Server) ListUsersAddresses(w http.ResponseWriter, req *http.Request) {
	var request struct {
		UserUuid string `json:"user_uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	addresses, err := s.service.FindAddressesByUserUuid(request.UserUuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, addresses)
	return
}
