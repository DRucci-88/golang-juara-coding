package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("produk tidak ditemukan")
	ErrSKUDuplicate    = errors.New("sku produ sudah terdaftar di sistem")
)

// Entity Bisnis Murni
type Product struct {
	ID        uint
	SKU       string
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Kontrak Interface untuk Repository
type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uint) (*Product, error)
	GetBySKU(ctx context.Context, sku string) (*Product, error)
}

// Kontrak Interface untuk Usecase
type ProductUsecase interface {
	Create(ctx context.Context, sku string, name string, price float64, stock int) (*Product, error)
	GetByID(ctx context.Context, id uint) (*Product, error)
}
