package entity

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	Username string `json:"username"`
	Email	string `json:"email"`
	Fullname string `json:"full_name"`
	Role     string `json:"role"`
	ID       int64    `json:"id"`
	jwt.RegisteredClaims
}

type ResetPasswordClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}