package models

import (
	"time"

	"gorm.io/gorm"
)

type BlacklistedToken struct {
	gorm.Model
	TokenString string    `gorm:"unique;not null" json:"token_string"`
	ExpiresAt   time.Time `gorm:"not null" json:"expired_at"`
}
