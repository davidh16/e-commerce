package repository

import "e-commerce/models"

func (r repository) Create(user models.User) (*models.User, error) {
	result := r.Db().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r repository) FindUserByEmailAddress(emailAddress string) (*models.User, error) {
	var user models.User
	result := r.Db().First(&user).Where("email_address=?", emailAddress)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
