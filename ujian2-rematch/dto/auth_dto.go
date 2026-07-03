package dto

import (
	"time"
	"ujian2_rematch/model"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey = []byte("KunciRahasiaNegaraSangatRahasiaSekali")

type JWTClaims struct {
	UserID     uint           `json:"user_id"`
	EmployeeID *uint          `json:"employee_id,omitempty"`
	Email      string         `json:"email"`
	Role       model.UserRole `json:"role"` // "ADMIN" atau "USER"
	jwt.RegisteredClaims
}

type AuthContext struct {
	UserID         uint           `json:"user_id"`
	EmployeeID     *uint          `json:"employee_id"`
	Email          string         `json:"email"`
	Role           model.UserRole `json:"role"`
	Token          string         `json:"token"`
	TokenExpiredAt time.Time      `json:"token_expired_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
