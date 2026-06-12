package main

import (
	"errors"
	"fmt"
)

// 1. Definisikan Sentinel Errors untuk Validasi Data
var (
	ErrEmptyName    = errors.New("Invalid: nama produ tidak boleh kosong")
	ErrNegativeVal  = errors.New("invalid: harga atau stok tidak boleh bernilai negatif")
	ErrInsufficient = errors.New("Checkout: stok barang tidak mencukupi")
	ErrNotFound     = errors.New("Checkout: barang tidak ditemukan di sistem")

	ErrQtyInvalid    = errors.New("Invalid: kuantitas belanja krang atau sama dengan 0")
	ErrInvalidCoupon = errors.New("Checkout: kuantitas atau nominal tidak cocok")
	ErrCartEmpty     = errors.New("Invalid: keranjang kosong")
)

// 2. Struct Model
type Item struct {
	Name  string
	Price float64
	Stock int
}
type StoreInventory struct {
	Items []Item
}

// 3. Method untuk validasi & tambah barang (Pointer Receiver)
func (inv *StoreInventory) AddProduct(name string, price float64, stock int) error {
	if name == "" {
		return ErrEmptyName
	}
	if price < 0 || stock < 0 {
		return ErrNegativeVal
	}
	newItem := Item{Name: name, Price: price, Stock: stock}
	inv.Items = append(inv.Items, newItem)
	return nil
}

// 4. Method untuk pemrosesan penjualan (Pointer Receiver)
func (inv *StoreInventory) SellProduct(name string, qty int) (float64, error) {
	for i, item := range inv.Items {
		if item.Name == name {
			if item.Stock < qty {
				// Menggunakan Error Wrapping untuk menyisipkan info status
				return 0, fmt.Errorf("%w Stock Tersedia: %d, Permintaan: %d", ErrInsufficient, item.Stock, qty)
			}

			// Kurangi stok pada backing array asli
			inv.Items[i].Stock -= qty
			totalIncome := item.Price * float64(qty)
			return totalIncome, nil
		}
	}
	// Mengembalikan Sentinel Error
	return 0, fmt.Errorf("%w: produk '%s'", ErrNotFound, name)
}

func main() {
	store := StoreInventory{[]Item{Item{Name: "A", Price: 12.1, Stock: 10}}}
	fmt.Println(store)

	_ = store.AddProduct("Laptop Gaming", 15_000_000.00, 5)
	_ = store.AddProduct("Mouse Wireless", 35_000.00, 15)

	fmt.Println("=============")
	fmt.Println("Simulasi Transaksi E-Commerce")
	fmt.Println("=============")

	fmt.Println(store)

	// Uji 1: Penjualan Berhasil
	if total, err := store.SellProduct("Mouse Wireless", 5); err != nil {
		fmt.Println("Gagal Transaksi", err)
	} else {
		fmt.Printf("[SUKSES] Penjualan Mouse Wireless x5. Total: Rp %.2f\n", total)
	}

	// Uji 2: Penjualan Gagal (Stok Kurang) - Menguji Is dengan error wrapped
	_, errSales := store.SellProduct("Laptop Gaming", 10)
	if errSales != nil {
		fmt.Println("[GAGAL]", errSales)
		if errors.Is(errSales, ErrInsufficient) {
			fmt.Println("Tindakan: Memicu notifikasi otomatis ke tim gudang")
		}
	}

	// Uji 3: Penjualan Gagal (Barang Tidak Ditemukan)
	_, errNotFound := store.SellProduct("Keyboard RGB", 1)
	if errNotFound != nil {
		fmt.Println("[GAGAL]", errNotFound)
	}
	fmt.Println(store)

}
