package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
Engine Kalkulator Diskon Bertingkat & Membership Premium
*/

type Shop struct {
	Name     string
	Kategori string
	Harga    float64
}

// type Cart struct {
// 	Shop Shop
// 	qty  int
// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	member := readString(reader, "Masukan Member: ")

	items := []string{
		"Samsung M23 | Electronics | 3_000_000",
		"Iphone 12 | Electronics | 7_500_000",
		"Balenciaga | Fashion | 23_000_000",
		"Supreme | Fashion | 30_000_000",
		"Baju Pekalongan | Fashion | 300_000",
		"Celana Cipayung | Fashion | 300_000",
		"Apel (1KG) | Grocery | 40_000",
		"Jeruk (1KG) | Grocery | 35_000",
	}
	shops := []Shop{}
	carts := make(map[int]int)
	totalBelanja := 0.0

	for _, item := range items {
		parts := strings.Split(item, "|")
		name := strings.TrimSpace(parts[0])
		kategori := strings.TrimSpace(parts[1])
		harga, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		shop := Shop{Name: name, Kategori: kategori, Harga: harga}
		shops = append(shops, shop)
		// fmt.Println(shop)
	}

	for {
		fmt.Println("Menu 0 untuk Checkout")
		shopPrettyPrint(&shops)
		cartPrettyPrint(&carts, &shops, &totalBelanja)
		fmt.Println("\n=====\n")
		menu := readInt(reader, "Masukan Pilihan Menu: ")
		if menu == 0 {
			break
		}
		qty := readInt(reader, "Quantity: ")
		carts[menu-1] = qty
		clearConsole()
	}

	switch member {
	case "PREMIUM":
		potonganDiskon := totalBelanja * 0.05
		totalBelanja = totalBelanja - potonganDiskon
		fmt.Println("Member [PREMIUM] mendapatkan diskon 5%")
		fmt.Printf("Potongan Diskon: Rp %.2f\n", potonganDiskon)
	case "GOLD":
		potonganDiskon := totalBelanja * 0.1
		totalBelanja = totalBelanja - potonganDiskon
		fmt.Println("Member [GOLD] mendapatkan diskon 10%")
		fmt.Printf("Potongan Diskon: Rp %.2f\n", potonganDiskon)
	}

	if totalBelanja > 2_000_000 {
		fmt.Println("Selamat! Anda mendapatkan Gratis Ongkir")
	}
	fmt.Printf("Total Belanja: Rp %.2f", totalBelanja)
}

func shopPrettyPrint(shops *[]Shop) {
	fmt.Println("Produk List:")
	for idx, item := range *shops {
		fmt.Printf("%d. Nama: %s, Harga: Rp %.02f, Kategori: %s\n", idx+1, item.Name, item.Harga, item.Kategori)
	}
}

func cartPrettyPrint(carts *map[int]int, shops *[]Shop, totalBelanja *float64) {
	if carts == nil || len(*carts) == 0 {
		return
	}
	fmt.Print("\nProduk yang masuk ke keranjang (Cart List):")
	for key, qty := range *carts {
		diskon := 0.0
		item := (*shops)[key]
		fmt.Printf("\n- Name: %s, Qty: %d", item.Name, qty)
		switch {
		case item.Kategori == "Electronics" && item.Harga > 5_000_000:
			diskon = 10
		case item.Kategori == "Fashion" && item.Harga > 500_000:
			diskon = 20
		case item.Kategori == "Grocery" && qty > 5:
			diskon = 5
		}
		harga := float64(qty) * item.Harga * (100.0 - diskon) / 100.0
		fmt.Printf(", Harga: Rp. %.2f", harga)
		*totalBelanja += harga
	}
	fmt.Printf("\n\nTotal Belanja: Rp %.2f", *totalBelanja)
}

func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
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
