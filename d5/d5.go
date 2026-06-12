package main

import "fmt"

// 1. Definisikan Interface kontrak pembayaran
type PaymentGateway interface {
	Pay(amount float64) error
}

// 2. Struct BankTransfer mengimplementasikan PaymentGateway secara implisit
type BankTransfer struct {
	BankName string
}

func (b BankTransfer) Pay(amount float64) error {
	fmt.Printf("Membayar Rp %.2f lewat transfer Bank %s\n", amount, b.BankName)
	return nil
}
func (b BankTransfer) Ko() error {
	return nil
}

// 3. Struct EWallet mengimplementasikan PaymentGateway secara implisit
type EWallet struct {
	Provider string
}

func (e EWallet) Pay(amount float64) error {
	fmt.Printf("Membayar Rp %.2f menggunakan E-Wallet %s\n", amount, e.Provider)
	return nil
}

// 4. Fungsi polimorfisme yang menerima interface apa pun yang cocok
func CheckoutOrder(gateway PaymentGateway, total float64) {
	err := gateway.Pay(total)
	if err != nil {
		fmt.Println("Pembayaran gagal:", err)
		return
	}
	fmt.Println("Transaksi Sukses!")
}

func ProcessData(d any) {
	switch v := d.(type) {
	case string:
		fmt.Printf("Data adalah string: %s (panjang %d)\n", v, len(v))
	case int:
		fmt.Printf("Data adalah int: %d (dikali dua = %d)\n", v, v*2)
	default:
		fmt.Println("Tipe data tidak didukung")
	}
}

func main() {
	fmt.Println("===============================\nSesi 1:Map dan Interfaces\n===============================")

	//// Sesi 1: Struktur Data Dinamis: Map Deep Dive (45 Menit)

	/// 1. Deklarasi dan Inisialisasi Map
	// Cara 1: Menggunakan make (sangat direkomendasikan)
	productPrices := make(map[string]float64)
	productPrices["Laptop"] = 15_000_000.0
	productPrices["Mouse"] = 35_000.0

	// Cara 2: Menggunakan map literal
	categories := map[string]string{
		"ROG": "Laptop Gaming",
		"MX3": "Mouse Productivity",
	}

	fmt.Println(productPrices, categories)
	fmt.Println(productPrices["Laptop"], categories["ROG"])
	fmt.Println(categories["AAA"], productPrices["AAA"]) // Tidak error, tapi zero-value

	/// 2.Memeriksa Keberadaan Kunci (Comma OK Idiom)
	price, exists := productPrices["AAA"]
	if !exists {
		fmt.Println("Barang tidak ditemukan")
	} else {
		fmt.Printf("Harga : %.2f", price)
	}

	/// 3. Menghapus Kunci & Iterasi Map
	for key, value := range productPrices {
		fmt.Printf("Produk %s -> harga %.2f\n", key, value)
	}
	// Menghapus data
	delete(productPrices, "Mouse")
	fmt.Println("Setelah Map di hapus")
	for key, value := range productPrices {
		fmt.Printf("Produk %s -> harga %.2f", key, value)
	}

	////////////////////////////////////////////////////////////
	//// Sesi 2: Abstraksi Melalui Interface (45 Menit)
	fmt.Print("\n\n===============================\nesi 2:Abstraksi Melalui Interface\n===============================\n")

	/// 1. Implemetasi Implisit (Duck Typing)
	/* Di  Go,  Anda  tidak  perlu  menulis  keyword  implements   untuk  menghubungkan  sebuah  struct
	dengan interface. Jika suatu struct memiliki semua metode dengan nama, parameter, dan nilai balik
	yang  persis  didefinisikan  oleh  sebuah  interface,  maka  struct  tersebut  otomatis
	mengimplementasikannya secara implisit */

	/// 2. Studi Kasus Abstraksi Pembayaran
	paymentMethod1 := BankTransfer{BankName: "BCA"}
	paymentMethod2 := EWallet{Provider: "GoPay"}

	// Fungsi yang sama bisa menerima dua tipe data konkret yang berbeda
	CheckoutOrder(paymentMethod1, 500_000)
	CheckoutOrder(paymentMethod2, 1_200_000)

	////////////////////////////////////////////////////////////
	//// Sesi 3: Empty Interface ( interface{}  /  any ) & Type Assertion(45 Menit)
	fmt.Print("\n\n===============================\nSesi 3: Empty Interface ( interface{}  /  any ) & Type Assertion(45 Menit)\n===============================\n")

	/// 1. Apa itu Empty Interface?
	var data any // atau var data interface{}
	data = "Halo"
	data = true
	data = 42
	fmt.Println(data)

	/// 2. Type Assertion & Type Switch
	// A. Type Assertion
	val, ok := data.(float64) // Mencoba mengkonversi data kembali ke int
	if ok {
		result := val + 10.1
		fmt.Println("Hasil:", result)
	} else {
		fmt.Println("Data bukan integer")
	}
	// B. Type Switch (Cara Lebih Rapi untuk Multi-tipe)
	ProcessData("Golang")
	ProcessData(100)
	ProcessData('\n')
}
