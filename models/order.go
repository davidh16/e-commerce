package models

type Order struct {
	ID            uint64
	Items         []Product
	OrderTotal    float64
	ShippingAddr  Address
	PaymentMethod PaymentMethod
}
