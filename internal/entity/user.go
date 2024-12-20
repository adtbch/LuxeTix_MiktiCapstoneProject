package entity

import "time"

type User struct {
	ID           int64     `json:"id"`
	Fullname     string    `json:"fullname" gorm:"column:fullname"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	Gender       string    `json:"gender"`
	Verify_token string    `json:"verify"`
	Reset_token  string    `json:"reset"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
