package models

import "time"

type Client struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	KeyID     string    `json:"key_id"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
