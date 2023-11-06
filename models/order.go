package models

type Order struct {
	Uuid          string `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Items         []Product
	OrderTotal    float64
	ShippingAddr  Address
	PaymentMethod PaymentInfo
}
