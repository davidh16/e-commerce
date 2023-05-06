package services

import "e-commerce/models"

func (s Service) Create(user models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return s.repository.Create(user)
}
