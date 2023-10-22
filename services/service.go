package services

import (
	"e-commerce/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redis                 *redis.Client
	userRepository        repository.UserRepository
	categoryRepository    repository.CategoryRepository
	subcategoryRepository repository.SubcategoryRepository
	productRepository     repository.ProductRepository
	mediaRepository       repository.MediaRepository
}

func NewService(
	redis *redis.Client,
	userRepo repository.UserRepository,
	categoryRepo repository.CategoryRepository,
	subcategoryRepo repository.SubcategoryRepository,
	productRepo repository.ProductRepository,
	mediaRepo repository.MediaRepository,
) Service {
	return Service{
		redis:                 redis,
		userRepository:        userRepo,
		categoryRepository:    categoryRepo,
		subcategoryRepository: subcategoryRepo,
		productRepository:     productRepo,
		mediaRepository:       mediaRepo,
	}
}
