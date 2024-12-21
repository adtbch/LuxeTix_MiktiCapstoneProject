package dto

import "time"

type NotificationInput struct {
	UserID    int64     `json:"user_id" validate:"required"`
	
	Message   string    `json:"message" validate:"required"`
	Is_Read   bool      `json:"is_read"`
	Create_at time.Time `json:"create_at"`
}
