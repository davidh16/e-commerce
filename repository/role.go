package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r roleRepository) Db() *gorm.DB {
	return r.db
}

type RoleRepository interface {
	Db() *gorm.DB
	FindRoleByUuid(uuid string) (*models.Role, error)
}

func (r roleRepository) FindRoleByUuid(uuid string) (*models.Role, error) {
	var role models.Role
	result := r.Db().Where("uuid=?", uuid).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
