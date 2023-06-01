package services

import (
	"e-commerce/models"
)

func (s Service) Create(user models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	result := s.repository.Db().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
