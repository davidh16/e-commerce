package repository

import (
	"e-commerce/models"
	"gorm.io/gorm"
)

type paymentInfoRepository struct {
	db *gorm.DB
}

func NewPaymentInfoRepository(db *gorm.DB) *paymentInfoRepository {
	return &paymentInfoRepository{
		db: db,
	}
}

func (r paymentInfoRepository) Db() *gorm.DB {
	return r.db
}

type PaymentInfoRepository interface {
	Db() *gorm.DB
	FindPaymentInfoByUuid(uuid string) (*models.PaymentInfo, error)
	FindPaymentInfoByUserUuid(userUuid string) ([]models.PaymentInfo, error)
}

func (r paymentInfoRepository) FindPaymentInfoByUuid(uuid string) (*models.PaymentInfo, error) {
	var paymentInfo models.PaymentInfo
	result := r.Db().Where("uuid=?", uuid).First(&paymentInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &paymentInfo, nil
}

func (r paymentInfoRepository) FindPaymentInfoByUserUuid(userUuid string) ([]models.PaymentInfo, error) {
	var paymentInfo []models.PaymentInfo
	result := r.Db().Where("user_uuid=?", userUuid).Find(&paymentInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return paymentInfo, nil
}
