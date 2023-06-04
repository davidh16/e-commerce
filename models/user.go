package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Uuid            string       `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	EmailAddress    string       `json:"email_address"`
	Password        string       `json:"password"`
	ShippingAddress *string      `json:"shipping_address"`
	PaymentInfo     *PaymentInfo `json:"payment_info" gorm:"embedded"`
	AccountStatus   int          `json:"account_status"`
	CreatedAt       time.Time    `json:"created_at"`
	//OrderHistory      []string      `json:"order_history"`
	//Wishlist          []string      `json:"wishlist"`
	//ShoppingCart      []string      `json:"shopping_cart"`
}

type PaymentInfo struct {
	CardHolderName string
	CardNumber     string
	ExpiryMonth    uint8
	ExpiryYear     uint16
	CVV            string
}

func (u *User) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(ValidationRules, User{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

var ValidationRules = map[string]string{
	"Email":    "required,email",
	"Password": "required,min=8",
}

func (u *User) TableName() string {
	return "users"
}
