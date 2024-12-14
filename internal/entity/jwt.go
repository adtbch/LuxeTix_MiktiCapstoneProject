package entity

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	Username string `json:"username"`
	Fullname string `json:"full_name"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}
