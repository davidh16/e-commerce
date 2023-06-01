package repository

import (
	database "e-commerce/db"
	"gorm.io/gorm"
)

type Repository struct {
	database.BaseInterface
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		BaseInterface: database.NewBase(db),
	}
}
