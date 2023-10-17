package services

import (
	"e-commerce/models"
)

func (s Service) CreateProduct(category *models.Product) (*models.Product, error) {
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.productRepository.Db().Begin()
	result := tx.Create(category)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return category, nil
}

func (s Service) UpdateProduct(updatedProduct *models.Product) (*models.Product, error) {
	err := updatedProduct.Validate()
	if err != nil {
		return nil, err
	}

	tx := s.productRepository.Db().Begin()

	result := tx.Where("uuid=?", updatedProduct.Uuid).Save(&updatedProduct)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return updatedProduct, nil
}

func (s Service) DeleteProduct(uuid string) error {
	tx := s.productRepository.Db().Begin()

	result := tx.Where("uuid=?", uuid).Delete(models.Product{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}

func (s Service) GetProduct(uuid string) (*models.Product, error) {
	category, err := s.productRepository.FindProductByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s Service) ListProducts() ([]models.Product, error) {
	var categories []models.Product
	result := s.productRepository.Db().Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
