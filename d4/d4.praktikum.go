package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Struct AddressP
type AddressP struct {
	Street  string
	City    string
	ZipCode string
}

// 2. Struct VendorP
type VendorP struct {
	ID       string
	Name     string
	AddressP AddressP // Komposisi AddressP
}

// 3. Struct ProductP
type ProductP struct {
	SKU     string
	Name    string
	Price   float64
	Stock   int
	VendorP *VendorP // Menyimpan pointer ke VendorP (menghindari duplikasi objek vendor di memori)
}

// 4. Constructor Function untuk ProductP
func NewProduct(sku, name string, price float64, stock int, vendor *VendorP) (*ProductP, error) {
	if sku == "" || name == "" {
		return nil, errors.New("SKU dan nama produk wajib diisi")
	}
	if price <= 0 {
		return nil, errors.New("harga produk harus lebih dari nol")
	}
	if stock < 0 {
		return nil, errors.New("stok awal tidak boleh negatif")
	}
	if vendor == nil {
		return nil, errors.New("produk harus memiliki vendor yang valid")
	}

	return &ProductP{
		SKU:     sku,
		Name:    name,
		Price:   price,
		Stock:   stock,
		VendorP: vendor,
	}, nil
}

// 5. Method dengan Pointer Receiver untuk menambah stok
func (p *ProductP) Restock(qty int) error {
	if qty <= 0 {
		return errors.New("kuantitas restock harus lebih dari nol")
	}
	p.Stock += qty // Mengubah nilai asli stock
	return nil
}

// 6. Method dengan Pointer Receiver untuk memotong stok saat penjualan
func (p *ProductP) DeductStock(qty int) error {
	if qty <= 0 {
		return errors.New("kuantitas penjualan harus lebih dari nol")
	}
	if p.Stock < qty {
		return fmt.Errorf("stok tidak mencukupi untuk SKU %s (Stok: %d, Dibutuhkan: %d)", p.SKU, p.Stock, qty)
	}
	p.Stock -= qty // Mengubah nilai asli stock
	return nil
}

// 7. Method dengan Value Receiver untuk menampilkan rincian produk (Read-only)
func (p ProductP) PrintDetails() {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("SKU           : %s\n", p.SKU)
	fmt.Printf("Nama Produk   : %s\n", p.Name)
	fmt.Printf("Harga Satuan  : Rp %.2f\n", p.Price)
	fmt.Printf("Stok Saat Ini : %d unit\n", p.Stock)
	fmt.Printf("Dikirim Oleh  : %s (%s, %s)\n", p.VendorP.Name, p.VendorP.AddressP.Street, p.VendorP.AddressP.City)
	fmt.Println("--------------------------------------------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==================================================")
	fmt.Println("           SISTEM GUDANG LOGISTIK E-COMMERCE       ")
	fmt.Println("==================================================")

	// Inisialisasi VendorP & Alamat
	gudangUtama := &VendorP{
		ID:   "VND-XYZ",
		Name: "PT Logistik Bersama",
		AddressP: AddressP{
			Street:  "Kawasan Industri Blok B",
			City:    "Bekasi",
			ZipCode: "17111",
		},
	}

	fmt.Println("--- Input Data Produk Baru ---")
	sku := readString(reader, "Masukkan SKU Produk  : ")
	nama := readString(reader, "Masukkan Nama Produk : ")
	harga := readFloat(reader, "Masukkan Harga Satuan: ")
	stok := readInt(reader, "Masukkan Stok Awal   : ")
	fmt.Println("--------------------------------------------------")

	// Membuat produk baru melalui Constructor
	prod, err := NewProduct(sku, nama, harga, stok, gudangUtama)
	if err != nil {
		fmt.Println("Error Pembuatan Produk:", err)
		fmt.Println("\nTekan Enter untuk keluar...")
		_, _ = reader.ReadString('\n')
		return
	}

	// Tampilkan data awal
	fmt.Println("\nData Produk Baru Terdaftar:")
	prod.PrintDetails()

	// Loop utama (mirip while true)
	for {
		fmt.Println("\n=== MENU OPERASI GUDANG ===")
		fmt.Println("1. Restock Produk")
		fmt.Println("2. Penjualan (Potong Stok)")
		fmt.Println("3. Tampilkan Status Inventaris")
		fmt.Println("4. Keluar")
		fmt.Println("===========================")

		pilihan := readString(reader, "Pilih Menu (1-4): ")

		if pilihan == "4" {
			break
		}

		switch pilihan {
		case "1":
			fmt.Println("\n--- Simulasi Restock Gudang ---")
			qtyRestock := readInt(reader, "Masukkan jumlah restock masuk: ")
			fmt.Printf(">> Melakukan restock masuk sebanyak %d unit...\n", qtyRestock)
			if err := prod.Restock(qtyRestock); err != nil {
				fmt.Println("Gagal Restock:", err)
			} else {
				fmt.Printf("Stok Baru: %d unit\n", prod.Stock)
			}
		case "2":
			fmt.Println("\n--- Simulasi Penjualan ---")
			qtySell := readInt(reader, "Masukkan jumlah penjualan: ")
			fmt.Printf(">> Terjadi penjualan sebanyak %d unit...\n", qtySell)
			if err := prod.DeductStock(qtySell); err != nil {
				fmt.Println("Gagal Potong Stok:", err)
			}
		case "3":
			fmt.Println("\nStatus Akhir Inventaris Produk:")
			prod.PrintDetails()
		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan angka 1-4.")
		}
	}

	fmt.Println("\nKeluar dari program. Terima kasih!")
	fmt.Println("Tekan Enter untuk menutup aplikasi...")
	_, _ = reader.ReadString('\n')
}

// Helper untuk membaca input bertipe string secara bersih
func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}
