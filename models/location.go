package models

import "gorm.io/gorm"

type Location struct {
	Model    gorm.Model
	Name     string
	ParentID *uint
	Children []Location `gorm:"foreignkey:ParentID"`
}
