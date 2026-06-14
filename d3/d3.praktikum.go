package main

/*
Sesi 4: Praktikum Mandiri: "Sistem Manajemen Inventaris & Penjualan" (45 Menit)
*/

import (
	"errors"
	"fmt"
)

// 1. Definisikan Sentinel Errors untuk Validasi Data
var (
	ErrEmptyNameP    = errors.New("Invalid: nama produ tidak boleh kosong")
	ErrNegativeValP  = errors.New("invalid: harga atau stok tidak boleh bernilai negatif")
	ErrInsufficientP = errors.New("Checkout: stok barang tidak mencukupi")
	ErrNotFoundP     = errors.New("Checkout: barang tidak ditemukan di sistem")
)

// 2. Struct Model
type ItemP struct {
	Name  string
	Price float64
	Stock int
}
type StoreInventoryP struct {
	Items []ItemP
}

// 3. Method untuk validasi & tambah barang (Pointer Receiver)
func (inv *StoreInventoryP) AddProduct(name string, price float64, stock int) error {
	if name == "" {
		return ErrEmptyNameP
	}
	if price < 0 || stock < 0 {
		return ErrNegativeValP
	}
	newItem := ItemP{Name: name, Price: price, Stock: stock}
	inv.Items = append(inv.Items, newItem)
	return nil
}

// 4. Method untuk pemrosesan penjualan (Pointer Receiver)
func (inv *StoreInventoryP) SellProduct(name string, qty int) (float64, error) {
	for i, item := range inv.Items {
		if item.Name == name {
			if item.Stock < qty {
				// Menggunakan Error Wrapping untuk menyisipkan info status
				return 0, fmt.Errorf("%w Stock Tersedia: %d, Permintaan: %d", ErrInsufficientP, item.Stock, qty)
			}

			// Kurangi stok pada backing array asli
			inv.Items[i].Stock -= qty
			totalIncome := item.Price * float64(qty)
			return totalIncome, nil
		}
	}
	// Mengembalikan Sentinel Error
	return 0, fmt.Errorf("%w: produk '%s'", ErrNotFoundP, name)
}

func main() {
	store := StoreInventoryP{[]ItemP{ItemP{Name: "A", Price: 12.1, Stock: 10}}}
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
		if errors.Is(errSales, ErrInsufficientP) {
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
