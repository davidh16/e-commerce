package controller

import (
	"e-commerce/config"
	"e-commerce/services"
	"net/http"
)

type Controller struct {
	service *services.Service
	config  config.Config
}

func NewController(s *services.Service, cfg config.Config) Controller {
	return Controller{
		service: s,
		config:  cfg,
	}
}

func returnResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
