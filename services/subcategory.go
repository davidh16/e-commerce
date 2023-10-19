package services

import (
	"e-commerce/models"
)

func (s Service) CreateSubcategory(subcategory *models.Subcategory) (*models.Subcategory, error) {
	err := subcategory.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.subcategoryRepository.Db().Begin()
	result := tx.Create(subcategory)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return subcategory, nil
}

func (s Service) UpdateSubcategory(updatedSubcategory *models.Subcategory) (*models.Subcategory, error) {
	err := updatedSubcategory.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.subcategoryRepository.Db().Begin()

	result := tx.Where("uuid=?", updatedSubcategory.Uuid).Save(&updatedSubcategory)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return updatedSubcategory, nil
}

func (s Service) DeleteSubcategory(uuid string) error {
	tx := s.subcategoryRepository.Db().Begin()

	result := tx.Where("uuid=?", uuid).Delete(models.Subcategory{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}

func (s Service) GetSubcategory(uuid string) (*models.Subcategory, error) {
	subcategory, err := s.subcategoryRepository.FindSubcategoryByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return subcategory, nil
}

func (s Service) ListSubcategories() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory
	result := s.subcategoryRepository.Db().Find(&subcategories)
	if result.Error != nil {
		return nil, result.Error
	}
	return subcategories, nil
}
