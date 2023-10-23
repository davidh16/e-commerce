package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Address struct {
	Uuid          string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	StreetAddress string    `json:"street_address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	PostalCode    string    `json:"postal_code"`
	Country       string    `json:"country"`
	User          User      `json:"-" gorm:"foreignKey:UserUuid;references:Uuid"`
	UserUuid      string    `json:"user_uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *Address) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(AddressValidationRules, Address{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

var AddressValidationRules = map[string]string{
	"StreetAddress": "required",
	"City":          "required",
	"State":         "required",
	"PostalCode":    "required",
	"Country":       "required",
	"UserUuid":      "required,uuid",
}

func (u *Address) Table() string {
	return "addresses"
}
