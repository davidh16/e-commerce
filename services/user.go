package services

import (
	"e-commerce/models"
	"fmt"
)

func (s Service) Create(user models.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}
	err = s.repository.Create(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
