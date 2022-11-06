package main

import (
	"log"

	"example.com/m/models"
)

func main() {
	models.ConnectionDatabase()

	client := models.Client{
		Name:   "example.com",
		KeyID:  "123",
		Secret: "abc",
	}

	models.DB.Create(&client)
	log.Printf("Client created!")
	log.Printf("Client name: %s", client.Name)
	log.Printf("Client key: %s", client.KeyID)
	log.Printf("Client secret: %s", client.Secret)
}
