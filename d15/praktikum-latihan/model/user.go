package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"type=varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type=varchar(255);not null" json:"-"`
	Role      string    `gorm:"type=varchar(20);default:'USER'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
