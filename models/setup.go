package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	const (
		MYSQL_DATABASE = "godb"
	)

	dsn := dsn(MYSQL_DATABASE)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Masthead{})

	DB = database
}

func dsn(dbName string) string {
	const (
		MYSQL_PORT          = "3306"
		MYSQL_ROOT_PASSWORD = "root"
		MYSQL_USER          = "admin"
		MYSQL_PASSWORD      = "secret"
	)
	return fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_PORT, dbName)
}
