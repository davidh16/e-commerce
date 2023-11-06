package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateProduct(w http.ResponseWriter, req *http.Request) {
	var product models.Product

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.CreateProduct(&product)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusCreated, err, nil)
	return

}

func (s *Server) UpdateProduct(w http.ResponseWriter, req *http.Request) {

	var updatedProduct *models.Product

	err := json.NewDecoder(req.Body).Decode(&updatedProduct)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	exists, err := s.service.GetProduct(updatedProduct.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.UpdateProduct(exists.Merge(updatedProduct))
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	var product struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.GetProduct(product.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	err = s.service.DeleteProduct(product.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, nil)
	return
}

func (s *Server) GetProduct(w http.ResponseWriter, req *http.Request) {
	var product struct {
		Uuid string `json:"uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	productDb, err := s.service.GetProduct(product.Uuid)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, productDb)
	return
}

func (s *Server) ListProducts(w http.ResponseWriter, req *http.Request) {
	products, err := s.service.ListProducts()
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, products)
	return
}
