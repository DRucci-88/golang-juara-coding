package model

import "errors"

// Product bertipe Exported
type Product struct {
	SKU   string
	Name  string
	Price float64
	Stock int
}

// Method untuk validasi data (Exported)
func (p *Product) Validate() error {
	if p.SKU == "" || p.Name == "" {
		return errors.New("invalid product: SKU dan nama produk tidak boleh kosong")
	}
	if p.Price < 0 || p.Stock < 0 {
		return errors.New("invalid product: harga dan stok tidak boleh bernilai nol")
	}
	return nil
}
