package models

import "gorm.io/gorm"

type Catering struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(256);not null" json:"name"`
	IsAvailable bool    `gorm:"default:true" json:"is_available"`
	Price       float64 `gorm:"not null" json:"price"`
}
