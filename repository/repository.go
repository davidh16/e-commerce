package repository

import (
	database "e-commerce/db"
	"e-commerce/models"
	"gorm.io/gorm"
)

type repository struct {
	database.BaseInterface
}

func NewRepository(db *gorm.DB) repository {
	return repository{
		BaseInterface: database.NewBase(db),
	}
}

type Repository interface {
	database.BaseInterface
	FindUserByEmailAddress(emailAddress string) (*models.User, error)
	SaveToken(token string) error
	Create(user models.User) (*models.User, error)
}

func (r repository) SaveToken(token string) error {
	result := r.Db().Table("refresh_tokens").Create(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
