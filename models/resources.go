package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name                   string `gorm:"type:varchar(110);not null" json:"name"`
	IsResourceAvailability bool   `gorm:"default:true" json:"is_available"`
}
