package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Product struct {
	Uuid        string    `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Color       string    `json:"color"`
	Code        string    `json:"code"`
	SubCategory string    `json:"sub_category"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Merge(x *Product) *Product {

	if x.Name != "" {
		p.Name = x.Name
	}
	if x.Description != "" {
		p.Description = x.Description
	}
	if x.Price != 0 {
		p.Price = x.Price
	}
	if x.ImageURL != "" {
		p.ImageURL = x.ImageURL
	}
	if x.Color != "" {
		p.Color = x.Color
	}
	if x.Code != "" {
		p.Code = x.Code
	}

	p.UpdatedAt = time.Now()
	return p
}

func (p *Product) TableName() string {
	return "products"
}

var productValidationRules = map[string]string{
	"Name":        "required",
	"Description": "required",
	"Price":       "required",
	"ImageURL":    "required",
	"Color":       "required",
	"Code":        "required",
	"SubCategory": "required",
	"Category":    "required",
}

func (p *Product) Validate() error {
	v := validator.New()
	v.RegisterStructValidationMapRules(productValidationRules, Product{})
	err := v.Struct(p)
	if err != nil {
		return err
	}
	return nil
}
