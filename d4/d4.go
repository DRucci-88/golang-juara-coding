package main

import (
	"errors"
	"fmt"
)

/*
Pemrograman Berorientasi Objek ala Go: Menguasai Struct, Method, dan Optimasi Pointer
*/

type Customer struct {
	ID    string
	Name  string
	Email string
	Age   int8
}

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Vendor struct {
	ID      string
	Name    string
	Address // Anonymous Field (Struct Embedding)
}

// Struct Tag (Sangat Penting untuk JSON Backen
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
	Stock int     `json:"qty"`
}

type Cart struct {
	TotalQty   int32
	TotalPrice float64
}

type Wallet struct {
	Owner   string
	Balance float64
}

func main() {
	/// Deklarasi dan Inisialisasi Struct
	fmt.Println("-- Deklarasi dan Inisialisasi Struct")
	// Cara 1: Menggunakan key-value (Sangat direkomendasikan karena aman dari perubahan)
	c1 := Customer{ID: "C-001", Name: "Dika", Email: "dika@danamas", Age: 28}

	// Cara 2: Positional (Tidak direkomendasikan untuk struct besar)
	c2 := Customer{"C-002", "Budi", "budi@gmail.com", 30}

	fmt.Println(c1, c2)
	fmt.Println(c1.Name) // Mengakses field menggunakan tanda titik (.)
	fmt.Println(c2.Email)

	////////////////////////////////////////////////////////////////////
	/// Komposisi Struct (Embedded Struct / Composition)
	fmt.Println("Komposisi Struct (Embedded Struct / Composition)")
	v1 := Vendor{
		ID:   "VND-99",
		Name: "Gudang Utama",
		Address: Address{
			Street:  "Jalan Sudirman No. 12",
			City:    "Jakarta Selatan",
			ZipCode: "12190",
		},
	}

	// Promosi Field: kita bisa mengakses property Address langsung dari object Vendor
	fmt.Println("Kota Vendor:", v1.City) // Setara dengan: v.Address.City

	////////////////////////////////////////////////////////////////////
	/// Memahami Alamat Memori ( &) dan Dereferensi ( * )
	fmt.Println("-- Memahami Alamat Memori ( &) dan Dereferensi ( * )")
	age := 25
	var ptr *int = &age // ptr menyimpan alamat dari memori age

	fmt.Println("Nilai age:", age)         // Output 25
	fmt.Println("Nilai ptr:", *ptr)        // Output 25
	fmt.Println("Alamat memori age:", ptr) // Output alamat acak
	fmt.Println("Alamat memori ptr:", &ptr)

	*ptr = 30 // Mengubah nilai di alamat memori ptr
	fmt.Println("Mengubah nilai dari ptr")
	fmt.Println("Nilai age:", age) // Output 25
	fmt.Println("Nilai ptr:", *ptr)

	/// Pointer ke Struct
	fmt.Println("-- Pointer ke Struc")
	v2 := Vendor{ID: v1.ID, Name: v1.Name, Address: v1.Address}
	fmt.Println(v1, v2)
	v2.City = "Bekasi"
	v2.ID = "VND-90"
	fmt.Println(v1, v2)

	var ptrV2 *Vendor = &v2
	fmt.Println(ptrV2, &ptrV2, *ptrV2)
	fmt.Println(ptrV2.Name) // Otomatis direferensikan oleh Go menampilkan value nya (syntactic  sugar)

	/// Bahaya Pointer: Nil Pointer Dereference
	fmt.Println("-- Bahaya Pointer: Nil Pointer Dereference")
	var v3 *Vendor
	// fmt.Println(v3.Address) // Panic Golang

	if v3 != nil {
		fmt.Println(v3.City)
	} else {
		fmt.Println("Vendor belum diinisialisasi")
	}

	////////////////////////////////////////////////////////////////////
	/// Pointer Receiver vs Value Receiver dalam Struct
	fmt.Println("-- Pointer Receiver vs Value Receiver dalam Struct")

	myCart := Cart{TotalQty: 2, TotalPrice: 150_000}
	myCart.UpdatePriceWrong(200_000)
	fmt.Println("Salah:", myCart.TotalPrice) // Output tetap: 150_000

	myCart.UpdatePriceRight(200_000)
	fmt.Println("Benar:", myCart.TotalPrice)

	/// Konstruktor di Go (Constructor Function)
	fmt.Println("-- Konstruktor di Go (Constructor Function)")
	myWallet, err := NewWallet("Budi", 500_000)
	if err != nil {
		fmt.Println("Gagal", err)
		return
	}
	fmt.Printf("Dompet milik %s berhasil dibuat dengan saldo Rp %.2f", myWallet.Owner, myWallet.Balance)
}

// Value Receiver: Data dicopy, perubahan didalam method tidak mengubah objek asli
func (c Cart) UpdatePriceWrong(newPrice float64) {
	c.TotalPrice = newPrice
}

// Pointer Receiver: Memanipulasi alamat asli, perubahan bersifat permanen
func (c *Cart) UpdatePriceRight(newPrice float64) {
	c.TotalPrice = newPrice
}

// Constructor Function: Mengembalikan pointer (*Wallet) agar hemat memori
func NewWallet(owner string, initialBalance float64) (*Wallet, error) {
	if owner == "" {
		return nil, errors.New("pemilik dompet tidak boleh kosong")
	}
	if initialBalance < 0 {
		return nil, errors.New("saldo awal tidak boleh negatif")
	}
	return &Wallet{
		Owner:   owner,
		Balance: initialBalance,
	}, nil
}
