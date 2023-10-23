package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *mediaRepository {
	return &mediaRepository{
		db: db,
	}
}

func (r mediaRepository) Db() *gorm.DB {
	return r.db
}

type MediaRepository interface {
	Db() *gorm.DB
	FindMediaByUuid(uuid string) (*models.Media, error)
}

func (r mediaRepository) FindMediaByUuid(uuid string) (*models.Media, error) {
	var media models.Media
	result := r.Db().Where("uuid=?", uuid).First(&media)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media, nil
}
