package services

import (
	"e-commerce/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redis          *redis.Client
	userRepository repository.UserRepository
}

func NewService(
	redis *redis.Client,
	userRepo repository.UserRepository,
) Service {
	return Service{
		redis:          redis,
		userRepository: userRepo,
	}
}
