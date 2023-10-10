package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) Db() *gorm.DB {
	return r.db
}

type UserRepository interface {
	Db() *gorm.DB
	FindUserByEmailAddress(emailAddress string) (*models.User, error)
	SaveToken(token string) error
}

func (r userRepository) SaveToken(token string) error {
	table := r.Db().Table("refresh_tokens")

	result := table.Create(map[string]interface{}{
		"token": token})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
