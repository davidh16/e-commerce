package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type ecommerceRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *ecommerceRepository {
	return &ecommerceRepository{
		db: db,
	}
}

type EcommerceRepository interface {
	Create(models.User) (*models.User, error)
}

func (r *ecommerceRepository) Create(user models.User) (*models.User, error) {
	col := r.db.Table("test")
	result := col.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
