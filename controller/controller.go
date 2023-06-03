package controller

import (
	"e-commerce/services"
	"net/http"
)

type Controller struct {
	service *services.Service
}

func NewController(s *services.Service) Controller {
	return Controller{
		service: s,
	}
}

func returnResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
