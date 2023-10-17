package services

import (
	"e-commerce/models"
)

func (s Service) CreateCategory(category *models.Category) (*models.Category, error) {
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.categoryRepository.Db().Begin()
	result := tx.Create(category)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return category, nil
}

func (s Service) UpdateCategory(updatedCategory *models.Category) (*models.Category, error) {
	err := updatedCategory.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.categoryRepository.Db().Begin()

	result := tx.Where("uuid=?", updatedCategory.Uuid).Save(&updatedCategory)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return updatedCategory, nil
}

func (s Service) DeleteCategory(uuid string) error {
	tx := s.categoryRepository.Db().Begin()

	result := tx.Where("uuid=?", uuid).Delete(models.Category{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}

func (s Service) GetCategory(uuid string) (*models.Category, error) {
	category, err := s.categoryRepository.FindCategoryByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s Service) ListCategories() ([]models.Category, error) {
	var categories []models.Category
	result := s.categoryRepository.Db().Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
