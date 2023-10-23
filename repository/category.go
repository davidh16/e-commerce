package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r categoryRepository) Db() *gorm.DB {
	return r.db
}

type CategoryRepository interface {
	Db() *gorm.DB
	FindCategoryByUuid(uuid string) (*models.Category, error)
}

func (r categoryRepository) FindCategoryByUuid(uuid string) (*models.Category, error) {
	var category models.Category
	result := r.Db().Where("uuid=?", uuid).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
