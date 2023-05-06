package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Uuid              string        `json:"id"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	Password          string        `json:"password"`
	ShippingAddressID string        `json:"shipping_address"`
	PaymentInfoID     PaymentMethod `json:"payment_info"`
	OrderHistory      []string      `json:"order_history"`
	Wishlist          []string      `json:"wishlist"`
	ShoppingCart      []string      `json:"shopping_cart"`
	AccountStatus     string        `json:"account_status"`
	CreatedAt         time.Time     `json:"created_at"`
}

type PaymentMethod struct {
	CardholderName string
	CardNumber     string
	ExpiryMonth    uint8
	ExpiryYear     uint16
	CVV            string
}

func (u User) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(ValidationRules, User{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

var ValidationRules = map[string]string{
	"Username": "required",
	"Email":    "required,email",
	"Password": "required,min=8",
}
