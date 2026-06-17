package main

/*
Praktikum Arsitektur Proyek Go: Membangun Backend Aplikasi E-Commerce Terstruktur & Multi-Library
*/

import (
	"fmt"
	"go-shop/auth"
	"go-shop/logger"
	"go-shop/model"
	"go-shop/service"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load berkas .env dari root direktori
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("[WARNING] Gagal memuat berkas .env, menggunakan variabel default")
	}

	// Membaca variable env
	port := os.Getenv("APP_PORT")
	envMode := os.Getenv("APP_ENV")
	secretToken := os.Getenv("MERCHANT_SECRET_TOKEN")

	fmt.Println("\n==================================================")
	fmt.Printf("   MEMULAI SERVER GO-SHOP DI PORT %s (%s) SecretKey [%s]\n", port, envMode, secretToken)
	fmt.Println("==================================================")

	// 2. Setup database invetaris dummy
	inventory := map[string]model.Product{
		"SKU-ROG-01": {SKU: "SKU-ROG-01", Name: "Laptop ROG Zephyrus", Price: 25_000_000},
		"SKU-MOU-02": {SKU: "SKU-MOU-02", Name: "Mouse Wireless Logitech", Price: 50_000},
	}

	// 3. Simulasi Transaksi 1: Authentikasi Gagal
	fmt.Println("\n[Transaksi 1] Kasir mencoba checkout dengan token tidak valid")
	authenticator := auth.NewAuthenticator(secretToken)

	tokenKasirSalah := "token-palsu-123"
	if !authenticator.ValidateToken(tokenKasirSalah) {
		logger.LogFailure("Akses ditolak: Token merchant kasir tidak cocok")
	}

	// 4. Simulasi Transaksi 2: Belanja Sukses
	fmt.Println("\n[TRANSAKSI 2] Kasir melakukan checkout belanja dengan token valid")
	tokenKasirBenar := "rahasia-kasir-goshop-2026"

	if authenticator.ValidateToken(tokenKasirBenar) {
		cart := service.NewCartService()

		// Tambahkan laptop ROG x2 dan mouse x3 ke keranjang
		_ = cart.AddItem(inventory["SKU-ROG-01"], 2)
		_ = cart.AddItem(inventory["SKU-MOU-02"], 3)
		// Lakukan checkout
		orderID, totalPay, errCheckout := cart.Checkout(inventory)
		if errCheckout != nil {
			logger.LogFailure(errCheckout.Error())
		} else {
			// Kurangi stok database inventaris riil (simulasi)
			p1 := inventory["SKU-ROG-01"]
			p1.Stock -= 2
			inventory["SKU-ROG-01"] = p1

			p2 := inventory["SKU-MOU-02"]
			p2.Stock -= 3
			inventory["SKU-MOU-02"] = p2
			// Catat log sukses
			logger.LogSuccess(orderID, totalPay)
		}
	}

	// 5. Simulasi Transaksi 3: Gagal karena stok habis
	fmt.Println("\n[TRANSAKSI 3] Kasir mencoba menjual melebihi sisa stok...")
	// if authenticator.ValidateToken(tokenKasirBenar) {
	cart := service.NewCartService()
	// Sisa stok ROG tinggal 3 unit. Mencoba beli 4 unit.
	errStock := cart.AddItem(inventory["SKU-ROG-01"], 4)
	if errStock != nil {
		logger.LogFailure(errStock.Error())
	}

	fmt.Println("\n==================================================")
}
