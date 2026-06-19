package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	password = "secret45"    // Ganti dengan password Anda
	dbname   = "ecommercedb" // Ganti dengan nama database Anda
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Koneksi ke Database PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Berhasil terhubung ke database!")

	// 2. Data Master Produk Riil berdasarkan Kategori (Masing-masing 10 produk)
	realProducts := map[string][]string{
		"Elektronik": {
			"Smartphone iPhone 15 Pro", "Samsung Galaxy S24 Ultra", "Laptop ASUS ROG Zephyrus",
			"TWS Sony WF-1000XM5", "Smart TV LG OLED 55 Inch", "iPad Air M2",
			"Keyboard Mechanical Keychron K2", "Mouse Logistics MX Master 3S", "Monitor Dell UltraSharp 27",
			"PlayStation 5 Slim",
		},
		"Pakaian": {
			"Kemeja Flanel Uniqlo", "Jaket Denim Levi's 501", "Celana Chino Erigo",
			"Kaos Polos Cotton Combed 30s", "Sweater Hoodie H&M", "Sepatu Sneakers Nike Air Jordan",
			"Celana Cargo Tactical", "Rok Plisket Panjang", "Batik Pria Lengan Panjang",
			"Sepatu Pantofel Kulit",
		},
		"Makanan": {
			"Indomie Goreng Aceh", "Keripik Singkong Maicih", "Cokelat Silverqueen Almond",
			"Roti Tawar Sari Roti", "Susu UHT Ultra Milk 1L", "Kopi Kenangan Mantan Bottle",
			"Biskuit Oreo Vanilla", "Sereal Kellogg's Corn Flakes", "Samyang Buldak Noodles",
			"Selai Nutella 350g",
		},
		"Buku": {
			"Buku Atomic Habits - James Clear", "Buku Filosofi Teras - Henry Manampiring", "Novel Bumi - Tere Liye",
			"Buku Berani Tidak Disukai", "Buku Laut Bercerita - Leila S. Chudori", "Buku Rich Dad Poor Dad",
			"Novel Gadis Kretek", "Buku Sebuah Seni untuk Bersikap Bodo Amat", "Buku Psikologi Uang",
			"Buku Sapiens - Yuval Noah Harari",
		},
		"Otomotif": {
			"Oli Mesin Shell Helix HX8", "Helm Full Face KYT Vendetta", "Kampas Rem Brembo Original",
			"Ban Motor Maxxis Extramaxx", "Parfum Mobil Little Trees", "Lap Microfiber Pembersih",
			"Busi Denso Iridium", "Jas Hujan Axio Rubber", "Car Charger Anker Dual Port",
			"Cairan Pembersih Jamur Kaca",
		},
	}

	// 3. Proses Loop & Insert
	for catName, prodList := range realProducts {
		// Insert Kategori
		var categoryID int
		queryCat := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
		err := db.QueryRow(queryCat, catName).Scan(&categoryID)
		if err != nil {
			log.Printf("Gagal insert kategori %s (mungkin sudah ada): %v\n", catName, err)
			// Jika kategori sudah ada, kita skip atau ambil ID-nya (di sini kita skip demi kesederhanaan)
			continue
		}
		fmt.Printf("\nKategori [%s] berhasil dibuat dengan ID: %d\n", catName, categoryID)

		// Insert 10 Produk Riil yang ada di list
		for i, prodName := range prodList {
			// Membuat SKU acak yang rapi
			sku := fmt.Sprintf("SKU-%s-%d%d", faker.UUIDDigit()[:4], categoryID, i+1)

			// Generate harga & stok acak yang masuk akal
			price := float64((rand.Intn(50) + 1) * 15000) // kelipatan 15.000 agar rapi
			stock := rand.Intn(80) + 10

			queryProd := `INSERT INTO products (sku, category_id, name, price, stock) 
						  VALUES ($1, $2, $3, $4, $5)`

			_, err := db.Exec(queryProd, sku, categoryID, prodName, price, stock)
			if err != nil {
				log.Printf("  Gagal insert produk %s: %v\n", sku, err)
			} else {
				fmt.Printf("  -> Berhasil [%s]: %-40s | Rp %10.2f | Stok: %d\n", sku, prodName, price, stock)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}

	fmt.Println("Proses pengisian data riil selesai!")
}
