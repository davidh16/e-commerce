package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Category struct {
	Uuid      string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Category) Merge(x *Category) *Category {
	c.Title = x.Title
	c.UpdatedAt = time.Now()
	return c
}

func (c *Category) TableName() string {
	return "categories"
}

var categoryValidationRules = map[string]string{
	"Title": "required",
}

func (c *Category) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(categoryValidationRules, Category{})
	err := v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
