package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type verificationTokenRepository struct {
	db *gorm.DB
}

func NewVerificationTokenRepository(db *gorm.DB) *verificationTokenRepository {
	return &verificationTokenRepository{
		db: db,
	}
}

func (r verificationTokenRepository) Db() *gorm.DB {
	return r.db
}

type VerificationTokenRepository interface {
	Db() *gorm.DB
}

func (r verificationTokenRepository) FindTokenByUuid(uuid string) (*models.Token, error) {
	var verificationToken models.Token
	result := r.Db().Where("uuid=?", uuid).First(&verificationToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return &verificationToken, nil
}
