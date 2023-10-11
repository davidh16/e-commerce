package repository

import (
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
}
