package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Golang JuaraCoding")

	// variabel produk

	/*

		product_name string
		price int
		stock int
		is_active bool

	*/

	product_name := "NVIDIA RTX 4090"
	price := 30000000
	stock := 5
	is_active := true
	description := "High-end"
	seller_id := "019ea77a-8975-7963-bd0a-d27ec370667b"

	fmt.Println("Product Name:", product_name)
	fmt.Println("Price:", price)
	fmt.Println("Stock:", stock)
	fmt.Println("Is Active:", is_active)
	fmt.Println("Description:", description)
	fmt.Println("Seller ID:", seller_id)

	// Zero Value
	var company string
	var number int
	var is_verified bool

	fmt.Println(company)
	fmt.Println(number)
	fmt.Println(is_verified)

	// Operator
	a := 10
	b := 5

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// subtotal cart
	quantity := 3
	total := quantity * price
	fmt.Println("Total Cart:", total)

	// Augmented Assignment
	a += b
	fmt.Println("a += b =", a)
	a -= b
	fmt.Println("a -= b =", a)
	a *= b
	fmt.Println("a *= b =", a)
	a /= b
	fmt.Println("a /= b =", a)
	a %= b
	fmt.Println("a %= b =", a)

	// Comparison Operator
	fmt.Println("a == b =", a == b)
	fmt.Println("a != b =", a != b)
	fmt.Println("a > b =", a > b) // 5 > 5
	fmt.Println("a < b =", a < b) // 6 < 5
	fmt.Println("a >= b =", a >= b)
	fmt.Println("a <= b =", a <= b)

	// input from keyboard company
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan Nama Perusahaan: ")
	scanner.Scan()
	company = scanner.Text()
	fmt.Println("Nama Perusahaan: ", company)

	fmt.Println("Masukkan Jumlah Karyawan: ")
	var employee int
	fmt.Scan(&employee)
	fmt.Println("Jumlah Karyawan: ", employee)

}
