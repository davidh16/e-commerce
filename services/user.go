package services

import (
	"e-commerce/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

func (s Service) ValidateCredentials(user models.User) (*models.User, error) {

	// finding the user by email
	dbUser, err := s.repository.FindUserByEmailAddress(user.EmailAddress)
	if err != nil {
		return nil, err
	}

	//check if passwords match
	if !checkPasswordHash(user.Password, dbUser.Password) {
		return nil, errors.New("invalid credentials")
	}

	return dbUser, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s Service) Test() *models.User {
	return s.repository.Test()
}
