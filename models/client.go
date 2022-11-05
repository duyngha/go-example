package models

import "time"

type Client struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
