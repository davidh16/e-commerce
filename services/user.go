package services

import "e-commerce/models"

func (s Service) Create(user models.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}
	return s.repository.Create(user)
}
