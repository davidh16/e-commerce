package repository

import "e-commerce/models"

func (r *repository) Create(user models.User) (*models.User, error) {
	result := r.Db().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
