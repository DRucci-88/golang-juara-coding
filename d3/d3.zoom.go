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
	product := ProductZ{
		Name:  "Laptop",
		Price: 10000000,
	}
	// Memanggil Value Receiver
	product.Display()
	// Memanggil Pointer Receiver
	product.ApplyDiscount(1000000) // Diskon 1jt
	product.Display()              // Harga sekarang: 9jt

	// order status
	order := OrderStatusZ(Processed)
	fmt.Println("Status Order:", order)

	// check error
	err := CheckInventoryZ(0)
	// Pengecekan aman menggunakan errors.Is
	if errors.Is(err, ErrStockEmptyZ) {
		fmt.Println("Gagal: Hubungi vendor untuk re-stock!")
	}

	// custom error
	err = &PaymentErrorZ{
		TransactionID: "TXN12345",
		Code:          502,
		Message:       "Gateway Timeout",
	}
	if payErr := new(PaymentErrorZ); errors.As(err, &payErr) {
		fmt.Println("Detail Error:", payErr.Message)
		fmt.Println("Status Code:", payErr.Code)
		fmt.Println("Ref ID:", payErr.TransactionID)
	}

	// wrap error
	err = GetProductZ()
	fmt.Println(err) // Output: gagal memuat produk: inventory: barang tidak terdaftar

	// errors.Is tetap bisa mendeteksi Sentinel Error di dalam rantai wrapping!
	if errors.Is(err, ErrItemNotFoundZ) {
		fmt.Println("Detail: Ternyata barang memang tidak ada di database.")
	}

	SafeExecutionZ()
	fmt.Println("Program tetap berjalan aman setelah pemulihan.")

}

type ProductZ struct {
	Name  string
	Price float64
}

// Value Receiver: Hanya membaca data
func (p ProductZ) Display() {
	fmt.Printf("Produk: %s | Harga: Rp %.2f\n", p.Name, p.Price)
}

// Pointer Receiver: Mengubah nilai asli di memori
func (p *ProductZ) ApplyDiscount(discount float64) {
	p.Price -= discount
}

type OrderStatusZ int

const (
	PendingZ OrderStatusZ = iota
	ProcessedZ
	ShippedZ
)

// Menambahkan method pada tipe alias int
func (s OrderStatusZ) String() string {
	return [...]string{"PENDINGZ", "PROCESSEDZ", "SHIPPEDZ"}[s]
}

func calculatePrice(quantity int, price int) int {
	total := quantity * price
	return total
}

var (
	ErrStockEmptyZ   = errors.New("inventory: stok barang habis")
	ErrItemNotFoundZ = errors.New("inventory: barang tidak terdaftar")
)

func CheckInventoryZ(stock int) error {
	if stock == 0 {
		return ErrStockEmptyZ
	}
	return nil
}

type PaymentErrorZ struct {
	TransactionID string
	Code          int
	Message       string
}

// Implementasi interface error
func (e *PaymentErrorZ) Error() string {
	return fmt.Sprintf("[Payment Error %d] Transaksi %s Gagal: %s", e.Code, e.TransactionID, e.Message)
}

func DatabaseQueryZ() error {
	return ErrItemNotFoundZ // Sentinel error
}
func GetProductZ() error {
	err := DatabaseQueryZ()
	if err != nil {
		// Wrapping error dengan %w
		return fmt.Errorf("gagal memuat produk: %w", err)
	}
	return nil
}

func SafeExecutionZ() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Sistem berhasil pulih dari panic:", r)
		}
	}()

	fmt.Println("Memulai proses...")
	panic("koneksi database terputus total!") // Memicu panic
	fmt.Println("Baris ini tidak akan pernah dieksekusi.")
}
