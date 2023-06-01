package controller

import (
	"e-commerce/services"
)

type Controller struct {
	service *services.Service
}

func NewController(s *services.Service) Controller {
	return Controller{
		service: s,
	}
}
