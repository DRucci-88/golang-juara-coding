package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	SKU       string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"sku" binding:"required"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Price     float64        `gorm:"type:numeric(12,2);not null" json:"price" binding:"required,gt=0"`
	Stock     int            `gorm:"type:int;default:0" json:"stock" binding:"required,min=0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Disembunyikan dari JSON response
}
