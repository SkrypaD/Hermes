package services

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	RoleID int    `json:"role_id"`
	Role   string `json:"role"`
	Login  string `json:"login"`
	jwt.RegisteredClaims
}

type JWTService interface {
	// Attempts to generate token with provided claims.
	GenerateToken(claims CustomClaims, key string) (string, error)
	// Verifies validity of provided token.
	VerifyToken(token string, key string) (bool, error)
	// Attempts to retreive claims from provided token
	GetClaims(token string, key string) (*CustomClaims, error)
}

type PHashService interface {
	// Attempts to hash provided password.
	HashPassword(password string) (string, error)
	// Checks if provided password is the same as original.
	VerifyPassword(hashedOriginal string, candiate string) error
}
