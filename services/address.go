package services

import "e-commerce/models"

func (s Service) CreateAddress(address *models.Address) (*models.Address, error) {
	result := s.addressRepository.Db().Create(address)
	if result.Error != nil {
		return nil, result.Error
	}
	return address, nil
}

func (s Service) FindAddressesByUuid(uuid string) (*models.Address, error) {
	return s.addressRepository.FindAddressByUuid(uuid)
}

func (s Service) FindAddressesByUserUuid(userUuid string) ([]models.Address, error) {
	return s.addressRepository.FindAddressesByUserUuid(userUuid)
}
