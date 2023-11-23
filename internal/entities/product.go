package entities

import "github.com/jinzhu/gorm"

// Product model for database
type Product struct {
	gorm.Model
	Name        string
	Description string
	ImageLink   string
	Price       float64
	Rating      string
	Merchant    string
}
