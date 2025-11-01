package orderapi

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	PriceRub    int            `json:"prub"`
	Images      pq.StringArray `json:"images" gorm:"type:varchar(256)[]"`
}

func NewProduct(name, descr string, prub int) *Product {
	return &Product{
		Name:        name,
		Description: descr,
		PriceRub:    prub,
	}
}
