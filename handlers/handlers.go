package handlers

import "e-commerce/services"

type UserHandler struct {
	service services.Service
}

func NewUserHandler(service services.Service) *UserHandler {
	return &UserHandler{service: service}
}
