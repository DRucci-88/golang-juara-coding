package model

import "errors"

type Product struct {
	SKU   string
	Name  string
	Price float64
	Stock int
}

// Validate memeriksa kelayakan data produk
func (p *Product) Validate() error {
	if p.SKU == "" || p.Name == "" {
		return errors.New("SKU dan nama produk tidak boleh kosong")
	}
	if p.Price <= 0 {
		return errors.New("harga produk harus lebih dari nol")
	}
	if p.Stock < 0 {
		return errors.New("stok produk tidak boleh bernilai negatif")
	}
	return nil
}
