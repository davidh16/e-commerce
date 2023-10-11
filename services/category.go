package services

import (
	"context"
	"e-commerce/models"
)

func (s Service) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	result := s.userRepository.Db().Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
