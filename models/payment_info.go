package models

type PaymentInfo struct {
	CardHolderName string
	CardNumber     string
	ExpiryMonth    uint8
	ExpiryYear     uint16
	CVV            string
}
