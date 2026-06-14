package main

/*
Praktikum Mandiri: "Sistem Checkout Keranjang & Verifikasi Transaksi" (40 Menit)
*/

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 1. Data dummy item belanja (nama_produk | harga_satuan | jumlah_beli | status
	cartItems := []string{
		"Laptop Asus ROG | 25000000 | 1 | ready",
		"Mouse Wireless | 350000 | 2 | ready",
		"Smartphone Samsung | 12000000 | 1 | out_of_stock",
		"T-Shirt Basic | 150000 | 3 | ready"}
	fmt.Println("==================================================")
	fmt.Println("       MEMULAI PROSES CHECKOUT KERANJANG          ")
	fmt.Println("==================================================")

	var totalBelanja float64

	// Label untuk menghentikan seluruh checkout jika terjadi indikasi fraud
CheckoutProcessing:
	for idx, item := range cartItems {
		// Pecah data string pembatas " | "
		parts := strings.Split(item, " | ")
		productName := parts[0]

		// Parsing harga dan jumlah dari string
		var price float64
		var qty int
		fmt.Sscanf(parts[1], "%f", &price)
		fmt.Sscanf(parts[2], "%d", &qty)
		stockStatus := parts[3]

		fmt.Printf("\n[Item #%d] Memproses: %s x%d\n", idx+1, productName, qty)

		// Percabangan Switch-Case untuk memeriksa stok
		switch stockStatus {
		case "out_of_stock":
			fmt.Println(">> Status: STOK HABIS! Melewati item ini.")
			continue
		case "ready":
			subtotal := price * float64(qty)
			totalBelanja += subtotal
			fmt.Printf(">> Status: Ready. Subtotal: Rp %.2f\n", subtotal)

			// Logika Bisnis: Proteksi transaksi nilai tinggi
			if subtotal > 10_000_000 {
				fmt.Println(">> Peringatan: Transaksi Risiko Tinggi")
				fmt.Println(">> Memilai prosedur verifikasi OTP Keamanan")

				maxRetries := 3
				isOtpSuccess := false

				// Perulangan retry OTP
				for attempt := 1; attempt <= maxRetries; attempt++ {
					fmt.Printf("-> Percobaan OTP %d/%d: Mengirimkan OTP", attempt, maxRetries)

					// Simulasi: Percobaan 1 & 2 gagal
					if attempt == 3 {
						isOtpSuccess = true
					}

					// Jeda simulasi process
					time.Sleep(500 * time.Millisecond)

					if isOtpSuccess {
						fmt.Println("[OK] OTP Sukses")
						break
					} else {
						fmt.Println("[FAILED] OTP Gagal")
					}
				}

				// Jika verifikasi OTP gagal total setelah 3 kali mencoba
				if !isOtpSuccess {
					fmt.Println("Fraud Detected: Verifikasi gagal semua")
					fmt.Println("Membatalkan seluruh transaksi")
					totalBelanja = 0         // Reset nilai belanja
					break CheckoutProcessing // Batalkan seluruh checkout
				}
			}
		}
	}

	fmt.Println("\n==================================================")
	if totalBelanja > 0 {
		fmt.Printf("   TOTAL TRANSAKSI : Rp %15.2f\n", totalBelanja)
		fmt.Println("   STATUS CHECKOUT : BERHASIL!")
	} else {
		fmt.Println("   STATUS CHECKOUT : DIBATALKAN / TIDAK ADA ITEM VALID")
	}
	fmt.Println("==================================================")
}
