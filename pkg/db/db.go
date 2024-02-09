package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect - function for connecting to database
func Connect(dbName string) (*gorm.DB, error) {

	db, err := gorm.Open(
		sqlite.Open(dbName),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	return db, err
}
