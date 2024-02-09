package model

import (
	"gorm.io/gorm"
)

// User - user struct
type User struct {
	gorm.Model

	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
