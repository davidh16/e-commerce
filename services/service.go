package services

import (
	"e-commerce/repository"
)

type Service struct {
	repository repository.EcommerceRepository
}

func NewService(repo repository.EcommerceRepository) Service {
	return Service{
		repository: repo,
	}
}
