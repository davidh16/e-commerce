package services

import (
	"e-commerce/repository"
)

type Service struct {
	repository repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return Service{
		repository: repo,
	}
}
