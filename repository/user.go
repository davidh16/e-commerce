package repository

import (
	"e-commerce/models"
	"e-commerce/orm"
	"gorm.io/gorm"
)

type userRepository struct {
	orm.Postgres[models.User]
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		Postgres: orm.NewPostgres[models.User](db),
	}
}

type UserRepository interface {
	orm.Postgres[models.User]
}
