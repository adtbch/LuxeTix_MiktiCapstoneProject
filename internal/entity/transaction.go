package entity

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id" gorm:"column:UserID"`
	EventID   int64     `json:"event_id" gorm:"column:EventID"`
	Quantity  int64     `json:"quantity"`
	Total     int64     `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
