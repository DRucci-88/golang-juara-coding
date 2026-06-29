package domain

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrAuthPasswordHashFailed  = errors.New("Password failed to proceed hashing")
	ErrAuthRegisterFailed      = errors.New("Register Failed")
	ErrAuthLoginFailed         = errors.New("Email or Password incorrect ")
	ErrAuthTokenGenerateFailed = errors.New("Token JWT Generation Failed")
)

type RegisterInputDto struct {
	Email    string
	Password string
	Role     Role

	NIK      string
	FullName string
}

type LoginInputDto struct {
	Email    string
	Password string
}

type AuthUsecase interface {
	Register(context.Context, *RegisterInputDto) (*User, error)
	Login(context.Context, *LoginInputDto) (string, error)
	Logout(authContext *AuthContext) error
}

var JWTSecretKey = []byte("KunciRahasiaNegaraSangatRahasiaSekali")

type JWTClaims struct {
	UserID uint
	Email  string
	Role   string
	jwt.RegisteredClaims
}

type AuthContext struct {
	UserID         uint
	Email          string
	Role           string
	Token          string
	TokenExpiresAt time.Time
}
