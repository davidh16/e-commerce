package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Category struct {
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
}

func (u *Category) TableName() string {
	return "categories"
}

var categoryValidationRules = map[string]string{
	"Tittle": "required",
}

func (u *Category) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(categoryValidationRules, Category{})
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
