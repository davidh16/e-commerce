package services

import "e-commerce/models"

func (s Service) CreateRole(role *models.Role) error {
	tx := s.roleRepository.Db().Begin()
	result := tx.Create(role)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (s Service) DeleteRole(uuid string) error {
	tx := s.roleRepository.Db().Begin()
	result := tx.Where("uuid=?", uuid).Delete(models.Role{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (s Service) ListRoles() ([]models.Role, error) {
	var roles []models.Role
	result := s.roleRepository.Db().Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
