package models

import "time"

type Masthead struct {
	ID        uint      `json:"id"`
	ImageURL  string    `json:"image_url"`
	Link      string    `json:"link"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Order     int       `json:"order"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
