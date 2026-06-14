package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
Sistem Validasi Merchant E-Commerce dengan Custom Error Struct & errors.As
*/

type ValidationType string

const (
	HARGA_TIDAK_WAJAR ValidationType = "HARGA_TIDAK_WAJAR"
	STOK_KOSONG       ValidationType = "STOK_KOSONG"
)

type Product2 struct {
	ID       string
	Nama     string
	Kategori string
	Harga    float64
	Stok     int
}

type MerchantValidationError struct {
	MerchantID     string
	ProductID      string
	ValidationType ValidationType
	Message        string
}

func (err *MerchantValidationError) Error() string {
	return fmt.Sprintf("[MerchantValidationError] ValidationType [%s] MerchantID [%s] ProductID [%s] Message [%s]", err.ValidationType, err.MerchantID, err.ProductID, err.Message)
}

type MarketplacePlatform2 struct {
	Products map[string][]Product2
}

func (market *MarketplacePlatform2) RegisterMerchantProduct(
	merchantID string,
	prod Product2,
) error {
	if prod.Harga < 10_000 {
		return &MerchantValidationError{
			MerchantID:     merchantID,
			ProductID:      prod.ID,
			ValidationType: HARGA_TIDAK_WAJAR,
			Message:        "Registrasi gagal: Harga sangat tidak masuk akal",
		}
		// return fmt.Errorf("%w", &MerchantValidationError{
		// 	MerchantID:     merchantID,
		// 	ProductID:      prod.ID,
		// 	ValidationType: HARGA_TIDAK_WAJAR,
		// 	Message:        "Registrasi gagal: Harga sangat tidak masuk akal",
		// })
	}
	if prod.Stok <= 0 {
		return &MerchantValidationError{
			MerchantID:     merchantID,
			ProductID:      prod.ID,
			ValidationType: STOK_KOSONG,
			Message:        "Registrasi gagal: Stok kosong",
		}
	}
	if merchant := market.Products[merchantID]; merchant == nil {
		fmt.Println("Merchant Slice is Nil")
		merchant = make([]Product2, 5)
	} else if len(merchant) == 0 {
		fmt.Println("Merchant Slice is Empty")
		merchant = make([]Product2, 5)
	}
	market.Products[merchantID] = append(market.Products[merchantID], prod)
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("==================================================")
	fmt.Println("Sistem Validasi Merchant E-Commerce dengan Custom Error Struct & errors.As")
	fmt.Println("==================================================")

	market := MarketplacePlatform2{Products: make(map[string][]Product2)}
	var menu int
MainLoop:
	for {
		fmt.Println("\n0. Keluar aplikasi")
		fmt.Println("1. Registrasi Produk")
		fmt.Println("2. Show Produk")
		menu = readInt2(reader, "Masukan menu: ")

		switch menu {
		case 0:
			break MainLoop
		case 1:
			err := halamanRegistrasiProduk(reader, &market)
			showError(err)
		case 2:
			halamanShowProduct(&market)
		default:
			fmt.Println("Salah pilih menu")
		}
	}
	fmt.Println("Selamat Tinggal")
}

func halamanRegistrasiProduk(reader *bufio.Reader, market *MarketplacePlatform2) error {
	merchantId := readString2(reader, "Merchant ID: ")
	productID := readString2(reader, "Product ID: ")
	name := readString2(reader, "Name: ")
	kategori := readString2(reader, "Kategori: ")
	harga := readFloat2(reader, "Harga: ")
	stok := readInt2(reader, "Stok: ")

	// merchantId := "m1"
	// productID := "p1"
	// name := "n1"
	// kategori := "k1"
	// harga := 50_000.0
	// stok := 10

	product := Product2{
		ID:       productID,
		Nama:     name,
		Kategori: kategori,
		Harga:    harga,
		Stok:     stok,
	}
	fmt.Println(merchantId, product)

	return market.RegisterMerchantProduct(merchantId, product)
}

func showError(err error) {
	if err == nil {
		fmt.Println("Tidak ada error")
		return
	}
	fmt.Println(err)

	var merchantError *MerchantValidationError

	if errors.As(err, &merchantError) {
		fmt.Println(merchantError)
		errWrap := fmt.Errorf("MerchantID %s, ViolationType %s", merchantError.MerchantID, merchantError.ValidationType)
		fmt.Println(errWrap)
	} else {
		fmt.Println("GAWAT: ERROR TIDAK TERDETEKSI")
	}
}

func halamanShowProduct(market *MarketplacePlatform2) {
	fmt.Println("\nMenampilkan daftar barang per merchant")
	for merchantId, products := range market.Products {
		fmt.Printf("Merchant ID [%s]: \n", merchantId)
		for idx, product := range products {
			fmt.Printf("    %d. %v\n", idx, product)
		}
	}
	fmt.Println("\n===============\n")
}

// Helper untuk membaca input bertipe string secara bersih
func readString2(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt2(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString2(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat2(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString2(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}

func clearConsole2() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
