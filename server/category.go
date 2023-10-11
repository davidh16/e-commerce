package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateCategory(w http.ResponseWriter, req *http.Request) {
	var category *models.Category

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		returnResponse(req.Context(), w, http.StatusBadRequest, err)
		return
	}

	_, err = s.service.CreateCategory(req.Context(), category)

}

func (s *Server) UpdateCategory(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) DeleteCategory(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) GetCategory(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) ListCategories(w http.ResponseWriter, req *http.Request) {

}
