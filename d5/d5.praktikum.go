package main

import (
	"errors"
	"fmt"
)

/*
Praktikum Mandiri: "Sistem Multi-Payment Gateway & Cart Cache" (45 Menit)
*/

// 1. Interface Kontrak Pembayaran
type PaymentProcessor interface {
	Pay(amount float64) error
	GetMethodName() string
}

// 2. Concrete Struct 1: Midtrans CC
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) GetMethodName() string {
	return "Kartu Kredit (Midtrans)"
}

func (cc CreditCard) Pay(amount float64) error {
	if len(cc.CardNumber) < 16 {
		return errors.New("nomor kartu kredit tidak valid")
	}
	fmt.Printf("[PROSES] Memotong limit kartu kredit %s sebesar Rp %.2f\n", cc.CardNumber, amount)
	return nil
}

// 3. Concrete Struct 2: ShopeePay Qris
type QrisPayment struct {
	QrisID string
}

func (q QrisPayment) GetMethodName() string { return "QRIS (ShopeePay)" }

func (q QrisPayment) Pay(amount float64) error {
	if q.QrisID == "" {
		return errors.New("QRIS ID merchant kosong")
	}
	fmt.Printf("[PROSES] Menghasilkan QR Code dinamis untuk ID %s nominal Rp %.2f\n", q.QrisID, amount)
	return nil
}

// 4. Struct Keranjang Belanja menggunakan MAP (dynamic collection)
type Cart struct {
	CustomerID string
	Items      map[string]float64 // Map: [Nama Produk] -> Harga
}

// Constructor Cart
func NewCart(customerID string) *Cart {
	return &Cart{
		CustomerID: customerID,
		Items:      make(map[string]float64),
	}
}

// Method Tambah Item ke Map
func (c *Cart) AddItem(name string, price float64) {
	c.Items[name] = price
}

// Method hitung total & lakukan checkout dengan Payment Gateway interface
func (c *Cart) Checkout(processor PaymentProcessor) error {
	if len(c.Items) == 0 {
		return errors.New("checkout gagal: keranjang kosong")
	}
	var total float64
	fmt.Println("\nRincian Keranjang Belanja:")
	for item, price := range c.Items {
		fmt.Printf("- %-20s: Rp %12.2f\n", item, price)
		total += price
	}
	fmt.Printf("Total Tagihan        : Rp %12.2f\n", total)
	fmt.Printf("Metode Pembayaran    : %s\n", processor.GetMethodName())
	// Panggil interface method
	err := processor.Pay(total)
	if err != nil {
		return fmt.Errorf("pembayaran via %s gagal: %w", processor.GetMethodName, err)
	}
	// Transaksi sukses, kosongkan keranjang belanja
	c.Items = make(map[string]float64)
	return nil
}

func main() {
	// Inisialisasi keranjang belanja menggunakan map
	myCart := NewCart("CUST-DIKA-01")
	myCart.AddItem("Keyboard Mechanical", 1_200_000)
	myCart.AddItem("Mouse Wireless Office", 450_000)
	fmt.Println("==================================================")
	fmt.Println("          APLIKASI CHECKOUT MULTI-GATEWAY         ")
	fmt.Println("==================================================")
	// Uji 1: Sukses Bayar dengan QRIS
	qris := QrisPayment{QrisID: "QRIS-SHP-091"}
	errQris := myCart.Checkout(qris)
	if errQris != nil {
		fmt.Println("Transaksi Gagal:", errQris)
	} else {
		fmt.Println("Status Transaksi     : SUKSES!")
	}

	// Uji 2: Gagal Bayar karena data invalid (Nomor Kartu Kredit < 16 digit)
	myCart.AddItem("Monitor UltraWide", 4500000) // belanja lagi
	ccInvalid := CreditCard{CardNumber: "1234"}
	errCC := myCart.Checkout(ccInvalid)
	if errCC != nil {
		fmt.Println("\n[LOG KEAMANAN TRANSAKSI ERROR]")
		fmt.Println("Detail Kesalahan:", errCC)
	}
}
