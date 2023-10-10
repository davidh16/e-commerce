package services

import (
	"context"
	"e-commerce/config"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

func (s Service) GenerateJWT(uuid string, refreshToken bool) (string, error) {
	cfg := config.GetConfig()

	/*
		Checking if refreshToken is true, if it is, exp variable with value of 8 hours is set so refresh token can be build.
		If refreshToken is false, variable exp is set to 1 hour so access token with exp of 1 hour can be built.
	*/

	var exp int64

	if refreshToken {
		exp = time.Now().Add(8 * time.Hour).Unix()
	} else {
		exp = time.Now().Add(1 * time.Hour).Unix()
	}

	/*
		Generating new token. Worth knowing : SigningMethodHS256 allows secret key to be passed to SignedString method (line 64) as string.
		other methods require additional transformation of secret key to right type
	*/
	token := jwt.New(jwt.SigningMethodHS256)

	// storing some additional info in token as expiration, uuid of the user and if user is authorized
	token.Claims = jwt.MapClaims{
		"exp":        exp,
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

func (s Service) SaveRefreshToken(token string) error {
	return s.userRepository.SaveToken(token)
}

func (s Service) SaveAccessToken(userUuid string, token string) error {
	err := s.redis.Set(context.Background(), userUuid, token, 1*time.Hour).Err()
	if err != nil {
		return errors.New("Failed to store token in Redis")
	}
	return nil
}
