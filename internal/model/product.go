package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Name        string    `validate:"required"`
	Category    string    `validate:"required"`
	Description string    `validate:"required"`
	Price       float64   `validate:"required,number"`
	CreatedAt   time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:false"`
}
