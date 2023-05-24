package repository

import (
	database "e-commerce/db"
	"e-commerce/models"
	"gorm.io/gorm"
)

type repository struct {
	database.BaseInterface
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		BaseInterface: database.NewBase(db),
	}
}

type Repository interface {
	database.BaseInterface
	Create(user models.User) (*models.User, error)
}
