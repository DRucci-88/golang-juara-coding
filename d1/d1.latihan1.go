package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Menggunakan iota untuk mendefinisikan persentase diskon berdasarkan level member
const (
	DiscountNonMember1 float64 = 0.0  // 0%
	DiscountMember1    float64 = 0.05 // 5%
	DiscountPremium1   float64 = 0.10 // 10%
)

func main() {
	// Scanner untuk membaca input teks yang mengandung spasi
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("===========================================")
	fmt.Println("    APLIKASI KASIR TOKO - GOLANG BASE      ")
	fmt.Println("===========================================")

	// 1. Input Nama Pelanggan
	fmt.Print("Masukkan Nama Pelanggan : ")
	customerNameInput, _ := reader.ReadString('\n')
	customerName := strings.TrimSpace(customerNameInput)

	// 2. Input Nama Barang
	fmt.Print("Masukkan Nama Barang    : ")
	itemNameInput, _ := reader.ReadString('\n')
	itemName := strings.TrimSpace(itemNameInput)

	// 3. Input Harga Barang (dan konversi string => float64)
	fmt.Print("Masukkan Harga Barang   : ")
	priceInput, _ := reader.ReadString('\n')
	priceInput = strings.TrimSpace(priceInput)
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil || price <= 0 {
		fmt.Println("Error: Input harga barang harus berupa angka positif!")
		return
	}

	// 4. Input Jumlah Barang (dan konversi string => int)
	fmt.Print("Masukkan Jumlah Barang  : ")
	quantityInput, _ := reader.ReadString('\n')
	quantityInput = strings.TrimSpace(quantityInput)
	quantity, err := strconv.Atoi(quantityInput)
	if err != nil || quantity <= 0 {
		fmt.Println("Error: Input jumlah barang harus berupa angka bulat positif")
		return
	}

	// 5. Input Status Member (Non-Member / Member / Premium)
	fmt.Print("Status Member (none/member/premium): ")
	memberInput, _ := reader.ReadString('\n')
	memberStatus := strings.ToLower(strings.TrimSpace(memberInput))

	// //- PROSES KALKULASI & OPERASI DATA //-

	// Hitung subtotal kotor
	subtotal := price * float64(quantity)

	// Tentukan rate diskon awal berdasarkan status member
	var discountRate float64
	switch memberStatus {
	case "member":
		discountRate = DiscountMember1
	case "premium":
		discountRate = DiscountPremium1
	default:
		discountRate = DiscountNonMember1
	}
	// Logika Bisnis Tambahan (Operator Logika & Perbandingan):
	// Jika subtotal belanja lebih dari Rp 500.000, berikan tambahan diskon 2%
	// tanpa memandang status keanggotaan.
	hasExtraDiscount := subtotal > 500000.0
	if hasExtraDiscount {
		discountRate += 0.02
	}
	// Hitung nilai diskon
	discountAmount := subtotal * discountRate
	totalAfterDiscount := subtotal - discountAmount

	// Hitung PPN 11% dari total setelah diskon
	const taxRate = 0.11
	taxAmount := totalAfterDiscount * taxRate
	// Total Akhir yang harus dibayar
	grandTotal := totalAfterDiscount + taxAmount
	// //- FORMATTING OUTPUT (STRUK BELANJA) //-
	fmt.Println("\n===========================================")
	fmt.Println("               STRUK PEMBAYARAN            ")
	fmt.Println("===========================================")
	fmt.Printf("Pelanggan     : %s\n", customerName)
	fmt.Printf("Status Member : %s\n", strings.ToUpper(memberStatus))
	if hasExtraDiscount {
		fmt.Println("Promo         : Ekstra Diskon 2% (Belanja > 500k)")
	} else {
		fmt.Println("Promo         : Tidak ada promo tambahan")
	}
	fmt.Println("-------------------------------------------")
	fmt.Printf("%-20s x%-3d : Rp %12.2f\n", itemName, quantity, price)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Subtotal             : Rp %12.2f\n", subtotal)
	fmt.Printf("Diskon (%.0f%%)         : Rp %12.2f\n", discountRate*100, discountAmount)
	fmt.Printf("Total Setelah Diskon : Rp %12.2f\n", totalAfterDiscount)
	fmt.Printf("PPN (11%%)            : Rp %12.2f\n", taxAmount)
	fmt.Println("-------------------------------------------")
	fmt.Printf("TOTAL AKHIR          : Rp %12.2f\n", grandTotal)
	fmt.Println("===========================================")
	fmt.Println("       Terima Kasih Telah Berbelanja!      ")
	fmt.Println("===========================================")

	fmt.Print("Masukan uang tunai")
	cashPaidInput, _ := reader.ReadString('\n')
	cashPaidInput = strings.TrimSpace(cashPaidInput)
	cashPaid, err := strconv.ParseFloat(priceInput, 64)
	if err != nil || cashPaid <= 0.0 {
		fmt.Println("Error: Input uang tunai harus berupa angka positif!")
		return
	}

	if cashPaid < grandTotal {
		fmt.Println("Transaksi Gagal: Uang tunai tidak cukup!")
		return
	}

	kembalian := grandTotal - cashPaid
	fmt.Printf("Kembalian %.2f", kembalian)

	// priceInput, _ := reader.ReadString('\n')
	// priceInput = strings.TrimSpace(priceInput)
	// price, err := strconv.ParseFloat(priceInput, 64)
	// if err != nil || price <= 0 {
	// 	fmt.Println("Error: Input harga barang harus berupa angka positif!")
	// 	return
	// }
}
