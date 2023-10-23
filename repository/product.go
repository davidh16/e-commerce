package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) Db() *gorm.DB {
	return r.db
}

type ProductRepository interface {
	Db() *gorm.DB
	FindProductByUuid(uuid string) (*models.Product, error)
}

func (r productRepository) FindProductByUuid(uuid string) (*models.Product, error) {
	var product models.Product
	result := r.Db().Where("uuid=?", uuid).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
