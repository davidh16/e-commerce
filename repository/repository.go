package repository

import (
	database "e-commerce/db"
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
}
