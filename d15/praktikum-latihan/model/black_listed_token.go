package model

import "time"

type BlackListedToken struct {
	ID          uint      `gorm:"primaryKey"`
	TokenString string    `gorm:"type:text;uniqueIndex;not null"`
	ExpireAt    time.Time `gorm:"not null"`
}
