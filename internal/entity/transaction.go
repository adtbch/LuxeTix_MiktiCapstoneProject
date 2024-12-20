package entity

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id" gorm:"column:userid"`
	EventID   int64     `json:"event_id" gorm:"column:eventid"`
	Quantity  int64     `json:"quantity"`
	Amount    int64     `json:"amout"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
