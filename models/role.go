package models

import "github.com/go-playground/validator/v10"

type Role struct {
	Uuid string `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Name string `json:"name"`
}

func (m *Role) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(RoleValidationRules, Role{})
	err := v.Struct(m)
	if err != nil {
		return err
	}
	return nil
}

var RoleValidationRules = map[string]string{
	"Name": "required",
}

func (m *Role) TableName() string {
	return "roles"
}
