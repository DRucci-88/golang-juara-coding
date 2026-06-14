package day3latihan1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrEmptyName31     = errors.New("Invalid: nama produ tidak boleh kosong")
	ErrNegativeVal31   = errors.New("invalid: harga atau stok tidak boleh bernilai negatif")
	ErrInsufficient31  = errors.New("Checkout: stok barang tidak mencukupi")
	ErrNotFound31      = errors.New("Checkout: barang tidak ditemukan di sistem")
	ErrQtyInvalid31    = errors.New("Invalid: kuantitas belanja krang atau sama dengan 0")
	ErrInvalidCoupon31 = errors.New("Checkout: kuantitas atau nominal tidak cocok")
	ErrCartEmpty31     = errors.New("Invalid: keranjang kosong")
)

type Item31 struct {
	Name  string
	Price float64
	Stock int
}
type StoreInventory31 struct {
	Items []Item31
}

type CartItem31 struct {
	Name  string
	Price float64
	Qty   int
}
type ShoppingCart31 struct {
	Items []CartItem31
}

func (inv *StoreInventory31) AddInventory(name string, price float64, stock int) error {
	if name == "" {
		return ErrEmptyName31
	}
	if price < 0 || stock < 0 {
		return ErrNegativeVal31
	}
	newItem := Item31{Name: name, Price: price, Stock: stock}
	inv.Items = append(inv.Items, newItem)
	return nil
}

func (inv *StoreInventory31) SellProduct(name string, qty int) (float64, error) {
	for i, item := range inv.Items {
		if item.Name == name {
			if item.Stock < qty {
				// Menggunakan Error Wrapping untuk menyisipkan info status
				return 0, fmt.Errorf("%w Stock Tersedia: %d, Permintaan: %d", ErrInsufficient31, item.Stock, qty)
			}

			// Kurangi stok pada backing array asli
			inv.Items[i].Stock -= qty
			totalIncome := item.Price * float64(qty)
			return totalIncome, nil
		}
	}
	// Mengembalikan Sentinel Error
	return 0, fmt.Errorf("%w: produk '%s'", ErrNotFound31, name)
}

func (cart *ShoppingCart31) AddProductToCart(item CartItem31) error {
	for _, cartItem := range cart.Items {
		if cartItem.Name == item.Name {
			cartItem.Qty += item.Qty
			return nil
		}
	}
	cart.Items = append(cart.Items, item)
	return nil
}
func (cart *ShoppingCart31) ApplyDiscountCoupon(couponCode string) (float64, error) {

	switch couponCode {
	case "DISKON50":
		return float64(50), nil
	case "CASHBACK10":
		return float64(10), nil
	default:
		return 0.0, fmt.Errorf("%w Tidak ada diskon %s", ErrInvalidCoupon31, couponCode)
	}
}

func (cart *ShoppingCart31) CalculateTotal(discountRate float64) (float64, error) {
	if len(cart.Items) <= 0 {
		return 0.0, fmt.Errorf("%w Keranjang kosong", ErrCartEmpty31)
	}
	var totalPrice float64
	for _, cartItem := range cart.Items {
		totalPrice += (cartItem.Price * discountRate / 100) * float64(cartItem.Qty)
	}
	return totalPrice, nil
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	storeInventory := &StoreInventory31{Items: []Item31{}}
	shoppingCart := &ShoppingCart31{Items: []CartItem31{}}

	fmt.Println("==================================================")
	fmt.Println("           SISTEM GUDANG LOGISTIK E-COMMERCE       ")
	fmt.Println("==================================================")

	var menuHalaman int
	for {
		fmt.Println("=== Halaman Utama ===")
		fmt.Println("1. Masuk Sebagai Operator Gudang")
		fmt.Println("2. Masuk Sebagai Customer")
		fmt.Println("3. Keluar Applikasi")
		menuHalaman = readInt(reader, "Pilih menu (1-3): ")

		switch menuHalaman {
		case 1:
			halamanOperator(reader, storeInventory)
		case 2:
			halamanCustomer(reader, shoppingCart)
		case 3:
			fmt.Println("Selamat Tinggal")
			return
		default:
			fmt.Println("Salah pilih menu")
		}

	}
}

func halamanOperator(reader *bufio.Reader, storeInventory *StoreInventory31) {
	var menuHalamanOperator int
	for {

		fmt.Println("=== Halaman Operator ===")

		fmt.Println("1. Add Product To Inventory")
		fmt.Println("2. Apply Discount")
		fmt.Println("3. Balik ke Halaman Utama")

		menuHalamanOperator = readInt(reader, "Pilih menu (1-3): ")
		switch menuHalamanOperator {
		case 1:
		case 2:
		case 3:
			fmt.Println("Balik ke Halaman Utama")
			return
		default:
			fmt.Println("Salah pilih menu")
		}
	}
}

func addProductToInventory(reader *bufio.Reader) {
	fmt.Println("=== Halaman Add Product To Inventory ===")
	name := readString(reader, "Masukan nama produk: ")
	// price := read
}

func halamanCustomer(reader *bufio.Reader, shoppingCart *ShoppingCart31) {
	var menuHalamanCustomer int
	for {

		fmt.Println("=== Halaman Customer ===")

		fmt.Println("1. Add Product")
		fmt.Println("2. Apply Discount Coupon")
		fmt.Println("3. Calculate Total")
		fmt.Println("4. Display Shooping Cart")
		fmt.Println("5. Balik ke Halaman Utama")

		menuHalamanCustomer = readInt(reader, "Pilih menu (1-5): ")
		switch menuHalamanCustomer {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Println("Balik ke Halaman Utama")
			return
		default:
			fmt.Println("Salah pilih menu")
		}
	}
}

func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input Int tidak valid")
	}
}
