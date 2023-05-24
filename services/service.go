package services

import (
	"e-commerce/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repository: repo,
	}
}
