package db

import (
	"gorm.io/gorm"
)

// Repository - repository layer struct
type Repository struct {
	DB *gorm.DB
}

// New - function for creating new repository
func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
