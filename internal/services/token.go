package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenService interface {
	GenerateAccessToken(ctx context.Context, claims entity.JWTCustomClaims) (string, error)
	GenerateResetPasswordToken(ctx context.Context, claims entity.ResetPasswordClaims) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (jwt.MapClaims, error)
	ExtractUserIDFromToken(ctx echo.Context) (int64, error)
}

type tokenService struct {
	secretKey string
}

func NewTokenService(secretKey string) TokenService {
	return &tokenService{secretKey}
}

func (s *tokenService) GenerateAccessToken(ctx context.Context, claims entity.JWTCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}

func (s *tokenService) GenerateResetPasswordToken(ctx context.Context, claims entity.ResetPasswordClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}

func (s *tokenService) ValidateToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	// Mengparse dan memverifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing sesuai dengan HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		// Kembalikan secret key untuk verifikasi
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Mengembalikan claims dari token jika valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token claims")
	}
}

// ExtractUserIDFromToken is a helper function to extract UserID from JWT token
func (s *tokenService) ExtractUserIDFromToken(ctx echo.Context) (int64, error) {
	// 1. Mendapatkan token dari header Authorization
	tokenString := ctx.Request().Header.Get("Authorization")
	if tokenString == "" {
		return 0, fmt.Errorf("Authorization token is required")
	}

	// 2. Menghilangkan prefix "Bearer " jika ada
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// 3. Dekode token untuk mendapatkan claims (termasuk UserID)
	claims, err := s.ValidateToken(ctx.Request().Context(), tokenString)
	if err != nil {
		return 0, fmt.Errorf("invalid token")
	}

	// 4. Mengambil UserID dari claims
	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user ID not found in token")
	}

	// Mengembalikan userID dalam tipe int64
	return int64(userID), nil
}
