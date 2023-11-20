package services

import (
	"e-commerce/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (s Service) CreateUser(user *models.User) (*models.User, *models.Token, error) {
	err := user.Validate()
	if err != nil {
		return nil, nil, err
	}

	tx := s.userRepository.Db().Begin()

	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, nil, result.Error
	}

	token, err := s.GenerateVerificationToken()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	tx.Commit()

	return user, token, nil
}

func (s Service) VerifyUser(user *models.User) (*models.User, error) {
	tx := s.userRepository.Db().Begin()

	user.IsActive = true

	result := tx.Where("uuid=?", user.Uuid).Save(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return user, nil
}

func (s Service) GetUser(uuid string) (*models.User, error) {
	var user models.User
	result := s.userRepository.Db().Where("uuid").First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
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

func (s Service) Me(token string) (string, error) {
	payload, _, err := new(jwt.Parser).ParseUnverified(strings.Split(token, "Bearer ")[1], jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	var me string
	if claims, ok := payload.Claims.(jwt.MapClaims); ok {
		me = fmt.Sprint(claims["sub"])
	}

	if me == "" {
		return "", errors.New("not logged in")
	}
	return me, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
