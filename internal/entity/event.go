package entity

import "time"

type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id" gorm:"column:UserID"` // Sesuaikan dengan nama kolom di database
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Time        string    `json:"time"`
	Date        string    `json:"date"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
