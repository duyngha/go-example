package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Masthead{})

	DB = database
}
