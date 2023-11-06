package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateRole(w http.ResponseWriter, req *http.Request) {
	var role models.Role
	err := json.NewDecoder(req.Body).Decode(&role)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}
	err = s.service.CreateRole(&role)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, nil)
	return
}

func (s *Server) DeleteRole(w http.ResponseWriter, req *http.Request) {
	var request struct {
		Uuid string `json:"uuid"`
	}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}
	err = s.service.DeleteRole(request.Uuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, nil)
	return
}

func (s *Server) ListRoles(w http.ResponseWriter, req *http.Request) {
	roles, err := s.service.ListRoles()
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, roles)
	return
}
