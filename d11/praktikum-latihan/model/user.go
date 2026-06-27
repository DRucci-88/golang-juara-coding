package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"type=varchar(100);uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
}
