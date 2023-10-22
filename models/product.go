package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Product struct {
	Uuid            string      `json:"uuid" gorm:"unique;type:uuid; column:uuid;default:uuid_generate_v4()"`
	Name            string      `json:"name"`
	Brand           string      `json:"brand"`
	Description     string      `json:"description"`
	Price           float32     `json:"price"`
	MediaUuid       string      `json:"media_uuid"`
	Color           string      `json:"color"`
	Code            string      `json:"code"`
	SubcategoryUuid string      `json:"subcategory_uuid"`
	Subcategory     Subcategory `json:"-" gorm:"foreignKey:SubcategoryUuid;references:Uuid"`
	CategoryUuid    string      `json:"category_uuid"`
	Category        Category    `json:"-" gorm:"foreignKey:CategoryUuid;references:Uuid"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
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
	if x.MediaUuid != "" {
		p.MediaUuid = x.MediaUuid
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
	"MediaUuid":   "required",
	"Color":       "required",
	"Code":        "required",
	"Subcategory": "required,uuid,nefield=Category",
	"Category":    "required,uuid,nefield=Subcategory",
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
