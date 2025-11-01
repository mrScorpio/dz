package orderapi

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceRub    int    `json:"prub"`
	Images      string `json:"images" gorm:"default:no_image"`
}

func NewProduct(name, descr string, prub int) *Product {
	return &Product{
		Name:        name,
		Description: descr,
		PriceRub:    prub,
	}
}
