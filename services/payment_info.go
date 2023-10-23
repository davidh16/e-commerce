package services

import "e-commerce/models"

func (s Service) CreatePaymentInfo(paymentInfo *models.PaymentInfo) (*models.PaymentInfo, error) {
	result := s.paymentInfoRepository.Db().Create(paymentInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return paymentInfo, nil
}

func (s Service) FindPaymentInfoByUuid(uuid string) (*models.PaymentInfo, error) {
	return s.paymentInfoRepository.FindPaymentInfoByUuid(uuid)
}

func (s Service) FindPaymentInfoByUserUuid(userUuid string) ([]models.PaymentInfo, error) {
	return s.paymentInfoRepository.FindPaymentInfoByUserUuid(userUuid)
}
