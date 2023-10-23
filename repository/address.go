package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *addressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r addressRepository) Db() *gorm.DB {
	return r.db
}

type AddressRepository interface {
	Db() *gorm.DB
	FindAddressByUuid(uuid string) (*models.Address, error)
	FindAddressesByUserUuid(userUuid string) ([]models.Address, error)
}

func (r addressRepository) FindAddressByUuid(uuid string) (*models.Address, error) {
	var address models.Address
	result := r.Db().Where("uuid=?", uuid).First(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

func (r addressRepository) FindAddressesByUserUuid(userUuid string) ([]models.Address, error) {
	var address []models.Address
	result := r.Db().Where("user_uuid=?", userUuid).Find(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return address, nil
}
