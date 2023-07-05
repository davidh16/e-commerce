package services

import (
	"e-commerce/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redis      *redis.Client
	repository repository.Repository
}

func NewService(redis *redis.Client, repo repository.Repository) Service {
	return Service{
		redis:      redis,
		repository: repo,
	}
}
