package model

import (
	"gorm.io/gorm"
)

// Class - class struct
type Class struct {
	gorm.Model

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
