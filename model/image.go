package model

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model

	FileName string `json:"filename"`
	Label    string `json:"label"`
	Name     string `json:"name"`
	ClassID  uint   `json:"class_id"`
}
