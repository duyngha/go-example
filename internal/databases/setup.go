package databases

import (
	"fmt"
	"log"
	"os"

	"example.com/m/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Masthead{})
	database.AutoMigrate(&models.Client{})

	DB = database
}

func dsn() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(mysql:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, port, dbName)
}
