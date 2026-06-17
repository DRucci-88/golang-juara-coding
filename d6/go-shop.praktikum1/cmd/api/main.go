package main

import (
	"fmt"
	"go-shop-praktikum1/internal/model"
	"go-shop-praktikum1/internal/service"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("      APLIKASI KASIR GO-SHOP (MODULAR LAYOUT)     ")
	fmt.Println("==================================================")

	// 1. Definisikan Produk
	laptop := model.Product{
		SKU:   "SKU-LAP-01",
		Name:  "Laptop Business Pro",
		Price: 12500000.0,
		Stock: 5,
	}
	mouse := model.Product{
		SKU:   "SKU-MOU-02",
		Name:  "Mouse Wireless Silent",
		Price: 250000.0,
		Stock: 20,
	}
	// 2. Inisialisasi Service Keranjang Belanja
	cart := service.NewCartService()

	// 3. Masukkan Barang ke Keranjang
	fmt.Println("[PROSES] Memasukkan Laptop x2 ke keranjang...")
	if err := cart.AddItem(laptop, 2); err != nil {
		fmt.Println("Gagal:", err)
	}
	fmt.Println("[PROSES] Memasukkan Mouse x5 ke keranjang...")
	if err := cart.AddItem(mouse, 5); err != nil {
		fmt.Println("Gagal:", err)
	}

	// 4. Simulasi Pembelian Melebihi Stok
	fmt.Println("[PROSES] Memasukkan Laptop x10 ke keranjang...")
	if err := cart.AddItem(laptop, 10); err != nil {
		fmt.Println("Gagal:", err) // Diharapkan mencetak stok tidak cukup
	}

	// 5. Tampilkan Keranjang Belanja Akhir
	fmt.Println("\nRincian Akhir Keranjang Belanja:")
	for sku, qty := range cart.Items {
		fmt.Printf("- SKU: %-12s | Kuantitas: %2d unit\n", sku, qty)
	}
	fmt.Println("==================================================")
}
