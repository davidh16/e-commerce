package services

import "e-commerce/models"

func (s Service) CreatePaymentInfo(paymentInfo *models.PaymentInfo) (*models.PaymentInfo, error) {
	tx := s.paymentInfoRepository.Db().Begin()
	result := tx.Create(paymentInfo)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return paymentInfo, nil
}

func (s Service) FindPaymentInfoByUuid(uuid string) (*models.PaymentInfo, error) {
	return s.paymentInfoRepository.FindPaymentInfoByUuid(uuid)
}

func (s Service) FindPaymentInfoByUserUuid(userUuid string) ([]models.PaymentInfo, error) {
	return s.paymentInfoRepository.FindPaymentInfoByUserUuid(userUuid)
}
