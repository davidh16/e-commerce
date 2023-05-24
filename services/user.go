package services

import (
	"e-commerce/models"
	"fmt"
)

func (s Service) Create(user models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	result := s.repository.Db().Create(user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &user, result.Error
}
