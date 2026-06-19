package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	password = "secret45"    // Ganti dengan password Anda
	dbname   = "ecommercedb" // Ganti dengan nama database Anda
)

// Struktur bantuan untuk menampung data produk dari DB
type Product struct {
	ID    int
	Price float64
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Koneksi ke Database
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

	// 2. Ambil semua User ID yang valid dari database
	var userIDs []int
	rowsUser, err := db.Query("SELECT id FROM users")
	if err != nil {
		log.Fatal("Gagal mengambil data users:", err)
	}
	defer rowsUser.Close()
	for rowsUser.Next() {
		var uid int
		rowsUser.Scan(&uid)
		userIDs = append(userIDs, uid)
	}

	if len(userIDs) == 0 {
		log.Fatal("Tabel users kosong! Isi data users terlebih dahulu.")
	}

	// 3. Ambil semua Product ID dan Harganya yang valid dari database
	var products []Product
	rowsProd, err := db.Query("SELECT id, price FROM products")
	if err != nil {
		log.Fatal("Gagal mengambil data products:", err)
	}
	defer rowsProd.Close()
	for rowsProd.Next() {
		var p Product
		rowsProd.Scan(&p.ID, &p.Price)
		products = append(products, p)
	}

	if len(products) < 4 {
		log.Fatal("Jumlah produk unik di database minimal harus ada 4 untuk mengisi detail order!")
	}

	fmt.Println("Memulai proses pembuatan 25 data orders...")
	fmt.Println("----------------------------------------------------------------------")

	// 4. Loop Pembuatan 25 Order
OrderLoop:
	for i := 1; i <= 25; i++ {
		tx, err := db.Begin()
		if err != nil {
			log.Printf("Gagal memulai transaksi ke-%d: %v\n", i, err)
			continue
		}

		// Pilih User acak dan buat Nomor Order unik
		randomUserID := userIDs[rand.Intn(len(userIDs))]
		orderNumber := fmt.Sprintf("INV/%s/%d%03d", time.Now().Format("20060102"), rand.Intn(90)+10, i)

		// A. Insert ke tabel `orders` mula-mula dengan total_amount = 0 (Akan di-update setelah kalkulasi item)
		var orderID int
		queryOrder := `INSERT INTO orders (order_number, user_id, total_amount) VALUES ($1, $2, 0) RETURNING id`
		err = tx.QueryRow(queryOrder, orderNumber, randomUserID).Scan(&orderID)
		if err != nil {
			tx.Rollback()
			log.Printf("Order ke-%d Gagal: %v\n", i, err)
			continue
		}

		// B. Pilih 4 produk acak yang UNIK untuk dimasukkan ke detail item (menghindari duplikasi Composite PK)
		chosenProducts := make(map[int]Product)
		for len(chosenProducts) < 4 {
			p := products[rand.Intn(len(products))]
			chosenProducts[p.ID] = p
		}

		var totalAmount float64

		// C. Loop Insert 4 item ke tabel `order_items`
		for prodID, prodInfo := range chosenProducts {
			qty := rand.Intn(5) + 1 // Kuantitas 1-5 pcs
			itemPrice := prodInfo.Price
			totalAmount += itemPrice * float64(qty)

			queryItem := `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
			_, err = tx.Exec(queryItem, orderID, prodID, qty, itemPrice)
			if err != nil {
				tx.Rollback()
				log.Printf("  -> Gagal memasukkan item untuk Order ID %d: %v\n", orderID, err)
				continue OrderLoop // Membatalkan order ini sepenuhnya dan lanjut ke order berikutnya
			}
		}

		// D. Update total_amount akhir pada tabel `orders` berdasarkan kalkulasi item belanjaan
		queryUpdateTotal := `UPDATE orders SET total_amount = $1 WHERE id = $2`
		_, err = tx.Exec(queryUpdateTotal, totalAmount, orderID)
		if err != nil {
			tx.Rollback()
			log.Printf("  -> Gagal update total_amount untuk Order ID %d: %v\n", orderID, err)
			continue
		}

		// Selesai dan simpan transaksi
		err = tx.Commit()
		if err != nil {
			log.Printf("Gagal commit order ke-%d: %v\n", i, err)
		} else {
			fmt.Printf("[Order %02d] ID: %-3d | No: %-18s | User ID: %-2d | Total: Rp %.2f\n",
				i, orderID, orderNumber, randomUserID, totalAmount)
		}

	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Proses pengisian 25 data orders & 100 items selesai!")
}
