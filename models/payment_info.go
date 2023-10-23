package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type PaymentInfo struct {
	Uuid           string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	CardHolderName string    `json:"card_holder_name"`
	CardNumber     string    `json:"card_number"`
	ExpiryMonth    uint8     `json:"expiry_month"`
	ExpiryYear     uint16    `json:"expiry_year"`
	CVV            string    `json:"-" gorm:"-"`
	User           User      `json:"-" gorm:"foreignKey:UserUuid;references:Uuid"`
	UserUuid       string    `json:"user_uuid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *PaymentInfo) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(PaymentInfoValidationRules, PaymentInfo{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

var PaymentInfoValidationRules = map[string]string{
	"CardHolderName": "required",
	"CardNumber":     "required",
	"ExpiryMonth":    "required",
	"ExpiryYear":     "required",
	"CVV":            "required",
	"UserUuid":       "required,uuid",
}

func (u *PaymentInfo) Table() string {
	return "payment_info"
}
