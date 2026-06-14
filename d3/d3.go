package main

import (
	"errors"
	"fmt"
)

/*
Struktur Data dan Keamanan Kode: Mengelola Array, Method, dan Robust Error Handling di Go
*/

// / 1. Value Receiver vs Pointer Receiver
type Product struct {
	Name  string
	Price float64
}

// // Sesi 2: Method dalam Go (45 Menit)
// Value Receiver: Hanya membaca data
func (p Product) Display() {
	fmt.Printf("Produk: %s | Harga: Rp %.2f\n", p.Name, p.Price)
}

// Pointer Receiver: Mengubah nilai asli di memori
func (p *Product) ApplyDiscount(discount float64) {
	p.Price -= discount
}

// / 2. Method pada Tipe Kustom Non-Struct
type OrderStatus int

const (
	Pending OrderStatus = iota
	Processed
	Shipped
)

// Menambahkan method pada tipe alias int
func (s OrderStatus) String() string {
	return [...]string{"PENDING", "PROCESSED", "SHIPPED"}[s]
}

// / 3. Nil Receivers (Fitur Unik Go)
func (p *Product) GetNameSafely() string {
	if p == nil {
		return "Produk Tidak Tersedia"
	}
	return p.Name
}

// // Sesi 3: Robust Error Handling (45 Menit)
// / 1. Sentinel Errors (errors.Is)
var (
	ErrStockEmpty   = errors.New("inventory: stok barang habis")
	ErrItemNotFound = errors.New("inventory: barang tidak ada")
)

func CheckInventory(stock int) error {
	if stock == 0 {
		return ErrStockEmpty
	}
	return nil
}

// / 2. Custom Error Struct (errors.As)
type PaymentError struct {
	TransactionId string
	Code          int
	Message       string
}

// Implementasi interface error
func (e *PaymentError) Error() string {
	return fmt.Sprintf("[Payment Error %d] Transaksi %s Gagal: %s", e.Code, e.TransactionId, e.Message)
}

// / 3. Error Wrapping ( %w )
func DatabaseQuery() error {
	return ErrItemNotFound // Sentinel Error
}

func GetProduct() error {
	err := DatabaseQuery()
	if err != nil {
		// Wrapping error dengan %w
		return fmt.Errorf("gagal memuat produk: %w", err)
	}
	return nil
}

// / 4. Defer, Panic, dan Recover
func SafeExecution() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Sistem berhasil pulih dari panic:", r)
		}
	}()
	fmt.Println("Memulai proses...")
	panic("koneksi database terputus total!") // Memicu panic
	fmt.Println("Baris ini tidak akan pernah dieksekusi.")
}

func main() {
	//// Sesi 1: Koleksi Data dengan Array dan Slice (45 Menit)
	fmt.Println("============")
	fmt.Println("Sesi 1: Koleksi Data dengan Array dan Slice (45 Menit)")
	fmt.Println("============")

	/// 1. Array (Fixed Size)
	var prices [3]float64
	prices[0] = 15000.0
	prices[1] = 25000.5
	// prices[3] = 10000.0 // Compile Error: out of bounds
	// Ukuran array tidak bisa diubah. [3]int dan [5]int adalah tipe data yang berbed

	/// 2. Slice (Dynamic Size & Under the Hood)
	// Membuat slice dengan make: len = 3, cap = 5
	scores := make([]int, 3, 5)
	scores[0], scores[1], scores[2] = 85, 90, 95
	fmt.Printf("Scores: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	// Menambahkan data melebihi len (3), tapi masih di bawah cap (5)
	scores = append(scores, 100)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	scores = append(scores, 110)
	scores = append(scores, 120)
	scores = append(scores, 130)
	scores = append(scores, 140)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	scores = append(scores, 150)
	scores = append(scores, 160)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))
	scores = append(scores, 160)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	//// Sesi 2: Method dalam Go (45 Menit)
	fmt.Println("\n============")
	fmt.Println("Sesi 2: Method dalam Go (45 Menit)")
	fmt.Println("============")

	var orderStatus OrderStatus = Pending
	fmt.Println("Order Status: ", orderStatus.String())

	//// Sesi 3: Robust Error Handling (45 Menit)
	fmt.Println("\n============")
	fmt.Println("Sesi 3: Robust Error Handling (45 Menit)")
	fmt.Println("============")

	/// 1. Sentinel Errors (errors.Is)
	fmt.Print("\n--- Sentinel Errors (errors.Is)\n")
	err := CheckInventory(0)
	// Pengecekan aman menggunakan errors.Is
	if errors.Is(err, ErrStockEmpty) {
		fmt.Println("Gagal: Hubungi vendor untuk re-stock!")
	}
	fmt.Println(err.Error())

	/// 2. Custom Error Struct (errors.As)
	fmt.Printf("\n--- Custom Error Struct (errors.As)\n")
	err = fmt.Errorf("Request Failed: %w",
		&PaymentError{
			TransactionId: "TX-001",
			Code:          404,
			Message:       "Produk tidak ditemukan",
		},
	)
	var payErr *PaymentError
	if errors.As(err, &payErr) {

		fmt.Println(err)
		fmt.Println(payErr)
		fmt.Printf(
			"Gagal Bayar Kode %d pada ID Transaksi %s\n",
			payErr.Code,
			payErr.TransactionId)
	}

	/// 3. Error Wrapping ( %w )
	fmt.Print("\n--- Error Wrapping\n")
	err = GetProduct()
	fmt.Println(err, " - ", err.Error())

	/// 4. Defer, Panic, dan Recover
	fmt.Print("\n--- Defer, Panic, dan Recover\n")
	SafeExecution()
	fmt.Println("Program tetap berjalan aman setelah pemulihan.")
}
