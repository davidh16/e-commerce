package services

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

func (s Service) CreateMedia(media *models.Media) (*gorm.DB, error) {
	tx := s.mediaRepository.Db().Begin()
	result := tx.Create(media)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return tx, nil
}

func (s Service) FindMediaByUuid(uuid string) (*models.Media, error) {
	media, err := s.mediaRepository.FindMediaByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return media, nil
}
