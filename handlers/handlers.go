package handlers

import "e-commerce/services"

type UserHandler struct {
	s services.Service
}

func NewUserHandler(service services.Service) *UserHandler {
	return &UserHandler{s: service}
}
