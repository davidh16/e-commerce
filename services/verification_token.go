package services

import (
	"e-commerce/models"
	"e-commerce/utils"
	"time"
)

func (s Service) GenerateVerificationToken() (*models.Token, error) {
	token := &models.Token{
		Token:     utils.GenerateToken(),
		IsUsed:    false,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	result := s.verificationTokenRepository.Db().Create(token)
	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}

func (s Service) UpdateVerificationToken(token *models.Token) (*models.Token, error) {
	tx := s.verificationTokenRepository.Db().Begin()

	token.IsUsed = true

	result := tx.Where("uuid=?", token.Uuid).Save(&token)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return token, nil
}

func (s Service) GetToken(tokenString string) (*models.Token, error) {
	var token models.Token
	result := s.verificationTokenRepository.Db().Where("token=?", tokenString).First(&token)
	if result.Error != nil {
		return nil, result.Error
	}
	return &token, nil
}
