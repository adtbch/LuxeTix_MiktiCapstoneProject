package entity

import (
	"time"
)

type Notification struct {
	ID         int64       `json:"id"`
	UserID     int64     `json:"user_id" gorm:"column:userid"`
	Message    string    `json:"message"`
	IsRead     bool      `json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}