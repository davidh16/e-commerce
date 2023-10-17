package services

import (
	"e-commerce/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) CreateUser(user *models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.userRepository.Db().Begin()

	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return user, nil
}

func (s Service) ValidateCredentials(user models.User) (*models.User, error) {

	// finding the user by email
	dbUser, err := s.userRepository.FindUserByEmailAddress(user.EmailAddress)
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
