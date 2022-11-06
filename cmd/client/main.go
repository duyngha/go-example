package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"time"

	"example.com/m/models"
	"github.com/google/uuid"
)

func main() {
	models.ConnectionDatabase()

	//TODO: Determine client name from request or something
	client := clientGenerator("client_1")

	result := models.DB.Create(&client)

	if result != nil {
		return
	}

	log.Printf("Client created!")
	log.Printf("Client name: %s", client.Name)
	log.Printf("Client key: %s", client.KeyID)
	log.Printf("Client secret: %s", client.Secret)
}

func clientGenerator(name string) (client models.Client) {
	salt := []byte("secret_of_salt" + time.Now().String())
	sha := sha256.Sum256(salt)

	// Không baoh được lưu sha như 1 string trực tiếp trong DB
	// Nếu cần hiển thị sha cho user thì chuyển sang dạng Hexadecimal
	// Nếu cần 1 dạng string cho sha thì solution có là dùng base64
	secret := base64.StdEncoding.EncodeToString(sha[:])

	client = models.Client{
		Name:   name,
		KeyID:  uuid.New().String(),
		Secret: secret,
	}

	return
}
