package main

import (
	"fmt"
	"math/rand"
)

/*
Simulasi Sistem Flash Sale & Limit Stock (Advanced Loop)
*/

type Kategory int

const (
	Electronics Kategory = iota
	Fashion
	Grocery
)

type Product struct {
	Name     string
	Kategory Kategory
	Harga    float64
	Stock    int
}

func main() {
	shops := []Product{
		Product{Name: "Samsung M23", Kategory: Electronics, Harga: 3_000_000.0, Stock: 15},
		Product{Name: "Iphone 12", Kategory: Electronics, Harga: 7_500_000.0, Stock: 10},
		Product{Name: "Vivo 13", Kategory: Electronics, Harga: 5_500_000.0, Stock: 10},
		Product{Name: "Balenciaga", Kategory: Fashion, Harga: 30_000_000.0, Stock: 10},
		Product{Name: "Supreme", Kategory: Fashion, Harga: 23_000_000.0, Stock: 10},
		Product{Name: "Baju Pekalongan", Kategory: Fashion, Harga: 300_000.0, Stock: 10},
		Product{Name: "Celana Cipayung", Kategory: Fashion, Harga: 400_000.0, Stock: 10},
		Product{Name: "Apel (1KG)", Kategory: Grocery, Harga: 40_000, Stock: 10},
		Product{Name: "Jeruk (1KG)", Kategory: Grocery, Harga: 30_000, Stock: 10},
		Product{Name: "Pisang (1Tandan)", Kategory: Grocery, Harga: 20_000, Stock: 10},
	}

	customers := []string{"Alex", "Brian", "Charlie", "Drucci", "Ethan", "Farrel", "Gabriel", "Henry", "Ivan", "Jason"}

	// Shuffle customers
	rand.Shuffle(len(customers), func(i int, j int) {
		customers[i], customers[j] = customers[j], customers[i]
	})

	product := shops[0]
	luckyCustomers := make(map[string]int)

	for _, name := range customers {
		timeResponse := rand.Intn(100) + 1
		if timeResponse > 80 {
			fmt.Printf("Gagal: Koneksi Timeout %s\n", name)
			continue
		}
		random := rand.Intn(3) + 1
		buy := random
		if random > product.Stock {
			buy = product.Stock
		}
		luckyCustomers[name] = buy
		product.Stock -= buy
		if product.Stock == 0 {
			fmt.Printf("Sold Out: %s kehabisan stock", product.Name)
			break
		}
	}
	fmt.Printf("\n\nCustomer Beruntung yang mendapatkan Flash Sale Produk %s\n", product.Name)
	for name, qty := range luckyCustomers {
		fmt.Printf("%s mendapatkan %s sejumlah %d\n", name, product.Name, qty)
	}
}
