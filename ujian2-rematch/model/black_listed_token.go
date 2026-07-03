package model

import (
	"errors"
	"time"
)

type BlackListedToken struct {
	ID          uint      `gorm:"primaryKey"`
	TokenString string    `gorm:"type:text;uniqueIndex;not null"`
	ExpiredAt   time.Time `gorm:"not null"`
}

var (
	ErrTokenIsBlackListed   = errors.New("Token is blacklisted")
	ErrTokenValidatedFailed = errors.New("Failed to validate token")
	ErrTokenIsExpired       = errors.New("Token Expired")
)
