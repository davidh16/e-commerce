package services

import (
	"e-commerce/config"
	"e-commerce/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s Service) Create(user models.User) (*models.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	// encrypt given password
	user.Password, err = encryptPassword(user.Password)
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

func (s Service) GenerateJWT(uuid string) (string, error) {
	cfg := config.GetConfig()

	/*
		Generating new token. Worth knowing : SigningMethodHS256 allows secret key to be passed to SignedString method (line 64) as string.
		other methods require additional transformation of secret key to right type
	*/
	token := jwt.New(jwt.SigningMethodHS256)

	// storing some additional info in token as expiration, uuid of the user and if user is authorized
	token.Claims = jwt.MapClaims{
		"exp":        time.Now().Add(1 * time.Hour).Unix(),
		"authorized": true,
		"sub":        uuid,
	}

	// converting token into string
	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func encryptPassword(password string) (string, error) {

	// Hash the password using the generated salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		return "", err
	}

	// Return the hashed password as a string
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
