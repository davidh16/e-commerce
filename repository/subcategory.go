package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type subcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) *subcategoryRepository {
	return &subcategoryRepository{
		db: db,
	}
}

func (r subcategoryRepository) Db() *gorm.DB {
	return r.db
}

type SubcategoryRepository interface {
	Db() *gorm.DB
	FindSubcategoryByUuid(uuid string) (*models.Subcategory, error)
}

func (r subcategoryRepository) FindSubcategoryByUuid(uuid string) (*models.Subcategory, error) {
	var subcategory models.Subcategory
	result := r.Db().Where("uuid=?", uuid).First(&subcategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subcategory, nil
}
