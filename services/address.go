package services

import "e-commerce/models"

func (s Service) CreateAddress(address *models.Address) (*models.Address, error) {
	tx := s.addressRepository.Db().Begin()
	result := tx.Create(address)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return address, nil
}

func (s Service) FindAddressesByUuid(uuid string) (*models.Address, error) {
	return s.addressRepository.FindAddressByUuid(uuid)
}

func (s Service) FindAddressesByUserUuid(userUuid string) ([]models.Address, error) {
	return s.addressRepository.FindAddressesByUserUuid(userUuid)
}
