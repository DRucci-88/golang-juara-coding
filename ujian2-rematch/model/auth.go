package model

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey = []byte("KunciRahasiaNegaraSangatRahasiaSekali")

type JWTClaims struct {
	UserID     uint     `json:"user_id"`
	EmployeeID *uint    `json:"employee_id,omitempty"`
	Email      string   `json:"email"`
	Role       UserRole `json:"role"` // "ADMIN" atau "USER"
	jwt.RegisteredClaims
}

var (
	ErrAuthUnauthorized = errors.New("Email or Password Wrong")
)
