package model

import "github.com/golang-jwt/jwt/v5"

var JWTSecretKey = []byte("KunciRahasiaNegaraSangatRahasiaSekali")

type JWTClaims struct {
	UserID uint `json:"user_id"`
	// CustomerID uint   `json:"customer_id,omitempty"`
	Email string `json:"email"`
	Role  string `json:"role"` // "ADMIN" atau "USER"
	jwt.RegisteredClaims
}
