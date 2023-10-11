package services

import (
	"e-commerce/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redis              *redis.Client
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
}

func NewService(
	redis *redis.Client,
	userRepo repository.UserRepository,
	categoryRepo repository.CategoryRepository,
) Service {
	return Service{
		redis:              redis,
		userRepository:     userRepo,
		categoryRepository: categoryRepo,
	}
}
