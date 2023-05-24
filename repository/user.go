package repository

import (
	"e-commerce/orm"
	"gorm.io/gorm"
)

type repository struct {
	orm.Postgres
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		Postgres: orm.NewPostgres(db),
	}
}

type Repository interface {
	orm.Postgres
}
