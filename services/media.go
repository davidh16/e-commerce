package services

import (
	"context"
	"e-commerce/models"
	"mime/multipart"
)

func (s Service) CreateAndUploadMedia(media *models.Media, file multipart.File) (*models.Media, error) {
	tx := s.mediaRepository.Db().Begin()
	result := tx.Create(media)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	_, err := s.UploadMediaToBucket(context.Background(), file, media.Path)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return media, nil
}

func (s Service) FindMediaByUuid(uuid string) (*models.Media, error) {
	media, err := s.mediaRepository.FindMediaByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return media, nil
}
