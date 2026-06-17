package service

import (
	"errors"
	"fmt"
	"go-shop-latihan1/internal/model"
)

type CartService struct {
	Items map[string]int // Map: SKU -> Kuantitas
}

// Constructor CartService
func NewCartService() *CartService {
	return &CartService{
		Items: make(map[string]int),
	}
}

// AddItem bertipe Exported untuk diakses dari cmd/api/main.go
func (c *CartService) AddItem(prod model.Product, qty int) error {
	// Panggil validaasi model
	if err := prod.Validate(); err != nil {
		return err
	}
	if qty <= 0 {
		return errors.New("Kuantitas belanja harus lebih dari nol")
	}
	if prod.Stock < qty {
		return fmt.Errorf("stock produk '%s' tidak cukup (Sisa: %d)", prod.Name, prod.Stock)
	}

	// Tambahkan kuantitas ke map keranjang
	c.Items[prod.SKU] += qty
	return nil
}
