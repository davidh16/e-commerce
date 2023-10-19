package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Subcategory struct {
	Uuid         string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Title        string    `json:"title"`
	CategoryUuid string    `json:"category_uuid"`
	Category     Category  `json:"-" gorm:"foreignKey:CategoryUuid;references:Uuid"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Subcategory) Merge(x *Subcategory) *Subcategory {
	c.Title = x.Title
	c.UpdatedAt = time.Now()
	return c
}

func (c *Subcategory) TableName() string {
	return "subcategories"
}

var subcategoryValidationRules = map[string]string{
	"Title":        "required",
	"CategoryUuid": "required,uuid4",
}

func (c *Subcategory) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(subcategoryValidationRules, Subcategory{})
	err := v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
