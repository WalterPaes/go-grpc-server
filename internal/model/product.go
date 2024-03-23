package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Category    string
	Description string
	Price       float64
}
