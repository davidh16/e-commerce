package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateCategory(w http.ResponseWriter, req *http.Request) {
	var category models.Category

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.CreateCategory(&category)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusCreated, err, nil)
	return

}

func (s *Server) UpdateCategory(w http.ResponseWriter, req *http.Request) {

	var updatedCategory *models.Category

	err := json.NewDecoder(req.Body).Decode(&updatedCategory)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	exists, err := s.service.GetCategory(updatedCategory.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.UpdateCategory(exists.Merge(updatedCategory))
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) DeleteCategory(w http.ResponseWriter, req *http.Request) {
	var category struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.GetCategory(category.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	err = s.service.DeleteCategory(category.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) GetCategory(w http.ResponseWriter, req *http.Request) {
	var category struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	categoryDb, err := s.service.GetCategory(category.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, categoryDb)
	return
}

func (s *Server) ListCategories(w http.ResponseWriter, req *http.Request) {
	categories, err := s.service.ListCategories()
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, categories)
	return
}
