package entity

import "time"

type Event struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id" gorm:"column:UserID"` // Sesuaikan dengan nama kolom di database
	Description   string    `json:"description"`
	Title         string    `json:"title"`
	Time          string    `json:"time"`
	Date          string    `json:"date"`
	Location      string    `json:"location"`
	StatusEvent   string    `json:"status event" gorm:"column:StatusEvent"`
	StatusRequest string    `json:"status request" gorm:"column:StatusRequest"`
	Price         int64     `json:"price"`
	Category      string    `json:"category"`
	Quantity      int64     `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
