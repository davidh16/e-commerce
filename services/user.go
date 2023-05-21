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
	_, err = s.repository.Create(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, err
}
