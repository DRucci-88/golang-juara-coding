package model

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SKU       string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"sku"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Price     float64   `gorm:"type:numeric(12,2);not null" json:"price"`
	Stock     int       `gorm:"type:int;not null;default=0" json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}
