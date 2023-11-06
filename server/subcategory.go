package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateSubcategory(w http.ResponseWriter, req *http.Request) {
	var subcategory models.Subcategory

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&subcategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.CreateSubcategory(&subcategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusCreated, err, nil)
	return

}

func (s *Server) UpdateSubcategory(w http.ResponseWriter, req *http.Request) {

	var updatedSubcategory *models.Subcategory

	err := json.NewDecoder(req.Body).Decode(&updatedSubcategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	exists, err := s.service.GetSubcategory(updatedSubcategory.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.UpdateSubcategory(exists.Merge(updatedSubcategory))
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) DeleteSubcategory(w http.ResponseWriter, req *http.Request) {
	var subcategory struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&subcategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.GetSubcategory(subcategory.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	err = s.service.DeleteSubcategory(subcategory.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) GetSubcategory(w http.ResponseWriter, req *http.Request) {
	var subcategory struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&subcategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	subcategoryDb, err := s.service.GetSubcategory(subcategory.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, subcategoryDb)
	return
}

func (s *Server) ListSubategories(w http.ResponseWriter, req *http.Request) {
	subcategories, err := s.service.ListSubcategories()
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, subcategories)
	return
}
