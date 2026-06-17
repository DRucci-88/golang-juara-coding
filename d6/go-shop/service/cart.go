package service

import (
	"errors"
	"fmt"
	"go-shop/model"

	"github.com/google/uuid"
)

type CartService struct {
	Items map[string]int // Map: SKU -> Kuantitas beli
}

// NewCartService menginisialisasi keranjang belanja baru
func NewCartService() *CartService {
	return &CartService{Items: make(map[string]int)}
}

// AddItem memasukan item belanja ke keranjang dengan validasi stok
func (c *CartService) AddItem(prod model.Product, qty int) error {
	if err := prod.Validate(); err != nil {
		return err
	}
	if qty <= 0 {
		return errors.New("kuantitas belanja harus minimal 1 unit")
	}
	if prod.Stock < qty {
		return fmt.Errorf("stok produk '%s' tidak cukup (Sisa: %d)", prod.Name, prod.Stock)
	}
	c.Items[prod.SKU] += qty
	return nil
}

// Checkout memproses total belanja dan menerbitkan Order ID UUID baru
func (c *CartService) Checkout(
	products map[string]model.Product,
) (string, float64, error) {

	if len(c.Items) == 0 {
		return "", 0.0, errors.New("keranjang belanja masih kosong")
	}

	var totalAmount float64
	for sku, qty := range c.Items {
		prod, exists := products[sku]
		if !exists {
			return "", 0.0, fmt.Errorf("produk dengan SKU %s tidak ditemukan", prod.SKU)
		}
		totalAmount += prod.Price * float64(qty)
	}

	// Generate UUID untik untuk Order ID
	orderID := uuid.New().String()

	// Kosongkan keranjang belanja setelah checkout sukses
	c.Items = make(map[string]int)

	return orderID, totalAmount, nil

}
