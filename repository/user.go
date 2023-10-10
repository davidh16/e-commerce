package repository

import "e-commerce/models"

func (r userRepository) Create(user models.User) (*models.User, error) {
	result := r.Db().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r userRepository) FindUserByEmailAddress(emailAddress string) (*models.User, error) {
	var user models.User
	result := r.Db().First(&user).Where("email_address=?", emailAddress)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r userRepository) Test() *models.User {
	var user models.User

	result := r.Db().First(&user).Where("email_address=?", "david@david.hr")
	if result.Error != nil {
		return nil
	}
	return &user
}
