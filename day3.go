package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Function & Parameter")
	fmt.Println("Calculate Price:", calculatePrice(5, 100000))
	fmt.Println("Calculate Price:", calculatePrice(3, 200000))
	fmt.Println("Calculate Price:", calculatePrice(1, 150000))

	var prices [3]float64
	prices[0] = 15000.0
	prices[1] = 25000.5
	// prices[3] = 10000.0 // Compile Error: out of bounds
	// Ukuran array tidak bisa diubah. [3]int dan [5]int adalah tipe data yang berbeda
	fmt.Println(prices[1])

	// Membuat slice dengan make: len = 3, cap = 5
	scores := make([]int, 3, 5)
	scores[0], scores[1], scores[2] = 85, 90, 95

	fmt.Printf("Scores: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	// Menambahkan data melebihi len (3), tapi masih di bawah cap (5)
	scores = append(scores, 100)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", scores, len(scores), cap(scores))

	// Membuat struct
	product := Product3{
		Name:  "Laptop",
		Price: 10000000,
	}
	// Memanggil Value Receiver
	product.Display()
	// Memanggil Pointer Receiver
	product.ApplyDiscount(1000000) // Diskon 1jt
	product.Display()              // Harga sekarang: 9jt

	// order status
	order := OrderStatus(Processed)
	fmt.Println("Status Order:", order)

	// check error
	err := CheckInventory(0)
	// Pengecekan aman menggunakan errors.Is
	if errors.Is(err, ErrStockEmpty) {
		fmt.Println("Gagal: Hubungi vendor untuk re-stock!")
	}

	// custom error
	err = &PaymentError{
		TransactionID: "TXN12345",
		Code:          502,
		Message:       "Gateway Timeout",
	}
	if payErr := new(PaymentError); errors.As(err, &payErr) {
		fmt.Println("Detail Error:", payErr.Message)
		fmt.Println("Status Code:", payErr.Code)
		fmt.Println("Ref ID:", payErr.TransactionID)
	}

	// wrap error
	err = GetProduct()
	fmt.Println(err) // Output: gagal memuat produk: inventory: barang tidak terdaftar

	// errors.Is tetap bisa mendeteksi Sentinel Error di dalam rantai wrapping!
	if errors.Is(err, ErrItemNotFound) {
		fmt.Println("Detail: Ternyata barang memang tidak ada di database.")
	}

	SafeExecution()
	fmt.Println("Program tetap berjalan aman setelah pemulihan.")

}

type Product3 struct {
	Name  string
	Price float64
}

// Value Receiver: Hanya membaca data
func (p Product3) Display() {
	fmt.Printf("Produk: %s | Harga: Rp %.2f\n", p.Name, p.Price)
}

// Pointer Receiver: Mengubah nilai asli di memori
func (p *Product3) ApplyDiscount(discount float64) {
	p.Price -= discount
}

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

func calculatePrice(quantity int, price int) int {
	total := quantity * price
	return total
}

var (
	ErrStockEmpty   = errors.New("inventory: stok barang habis")
	ErrItemNotFound = errors.New("inventory: barang tidak terdaftar")
)

func CheckInventory(stock int) error {
	if stock == 0 {
		return ErrStockEmpty
	}
	return nil
}

type PaymentError struct {
	TransactionID string
	Code          int
	Message       string
}

// Implementasi interface error
func (e *PaymentError) Error() string {
	return fmt.Sprintf("[Payment Error %d] Transaksi %s Gagal: %s", e.Code, e.TransactionID, e.Message)
}

func DatabaseQuery() error {
	return ErrItemNotFound // Sentinel error
}
func GetProduct() error {
	err := DatabaseQuery()
	if err != nil {
		// Wrapping error dengan %w
		return fmt.Errorf("gagal memuat produk: %w", err)
	}
	return nil
}

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
