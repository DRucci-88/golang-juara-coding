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
Latihan 1: Sistem Keranjang Belanja dengan Kode Kupon & Sentinel Error
*/

var (
	ErrQtyInvalid1    = errors.New("Error jika kuantitas belanja kurang dari atau sama dengan 0.")
	ErrInvalidCoupon1 = errors.New("Error jika kuantitas atau nominal diskon tidak cocok.")
	ErrCartEmpty1     = errors.New("Error jika kalkulasi total dilakukan pada keranjang kosong")
)

type CartItem1 struct {
	Nama      string
	Harga     float64
	Kuantitas int
}

type ShoppingCart1 struct {
	Items []CartItem1
}

func (cart *ShoppingCart1) AddItem(item CartItem1) error {
	if item.Kuantitas <= 0 {
		return ErrQtyInvalid1
	}
	for idx, _ := range cart.Items {
		if item := cart.Items[idx]; item.Nama == item.Nama {
			item.Kuantitas += item.Kuantitas
			return nil
		}
	}
	cart.Items = append(cart.Items, item)
	return nil
}

func (cart *ShoppingCart1) ApplyDiscountCoupon(couponCode string) (float64, error) {
	switch couponCode {
	case "DISKON50":
		return 50.0, nil
	case "CASHBACK10":
		return 10.0, nil
	}
	return 0.0, ErrInvalidCoupon1
}

func (cart *ShoppingCart1) CalculateTotal(discountRate float64) (float64, error) {
	if cart == nil || len(cart.Items) <= 0 {
		return 0.0, ErrCartEmpty1
	}
	var totalBelanja float64 = 0.0
	for _, item := range cart.Items {
		totalBelanja += item.Harga * float64(item.Kuantitas)
	}
	totalBelanja = totalBelanja * (100 - discountRate) / 100
	return totalBelanja, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==================================================")
	fmt.Println("Sistem Keranjang Belanja dengan Kode Kupon & Sentinel Error")
	fmt.Println("==================================================")

	cart := &ShoppingCart1{Items: []CartItem1{}}
	var discountRate float64

	var menu int
MainLoop:
	for {
		fmt.Println("\n0. Keluar aplikasi")
		fmt.Println("1. Add Item")
		fmt.Println("2. Show Cart")
		fmt.Println("3. Apply Discount Coupon")
		fmt.Println("4. Checkout Belanja")
		menu = readInt1(reader, "Masukan menu: ")

		switch menu {
		case 0:
			break MainLoop
		case 1:
			halamanAddItem(reader, cart)
		case 2:
			halamanShowCart(cart)
		case 3:
			discountRate = halamanApplyDiscount(reader, cart)
		case 4:
			halamanCalculateTotal(cart, discountRate)
		default:
			fmt.Println("Salah pilih menu")
		}
	}
	fmt.Println("Selamat Tinggal")
}

func halamanAddItem(reader *bufio.Reader, cart *ShoppingCart1) {
	fmt.Println("Add Item to Shopping Cart")
	nama := readString1(reader, "Nama Produk: ")
	harga := readFloat1(reader, "Masukan Harga (Rp): ")
	kuantitas := readInt1(reader, "Masukan kuantitas: ")
	err := cart.AddItem(CartItem1{
		Nama:      nama,
		Harga:     harga,
		Kuantitas: kuantitas,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Gagal menambah produk (%s) %w", nama, err))
	}
}

func halamanShowCart(cart *ShoppingCart1) {
	fmt.Println("Show Shopping Cart: ")
	for idx, item := range cart.Items {
		fmt.Printf("%d. Nama %s, Harga %.2f, Kuantitas %d\n", idx+1, item.Nama, item.Harga, item.Kuantitas)
	}
}

func halamanApplyDiscount(reader *bufio.Reader, cart *ShoppingCart1) float64 {
	kupon := readString1(reader, "Masukan Kupon ")
	discountRate, err := cart.ApplyDiscountCoupon(kupon)
	if err != nil {
		fmt.Println(fmt.Errorf("Gagal apply diskon (%s) %w", kupon, err))
		return 0.0
	}
	fmt.Printf("Selamat kamu mendapatkan kupon %d percent\n", int(discountRate))
	return discountRate
}

func halamanCalculateTotal(cart *ShoppingCart1, discountRate float64) {
	totalBelanja, err := cart.CalculateTotal(discountRate)
	if err != nil {
		fmt.Println(fmt.Errorf("Gagal Checkout %w", err))
		return
	}

	fmt.Printf("Total Belanja Rp %.2f", totalBelanja)
}

// Helper untuk membaca input bertipe string secara bersih
func readString1(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt1(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString1(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat1(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString1(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}

func clearConsole1() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
