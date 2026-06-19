package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
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
	fmt.Println("Memulai proses insert 20 users dummy Indonesia...")
	fmt.Println("--------------------------------------------------------------------------------")

	// 2. Pool Data Khas Indonesia
	firstNames := []string{
		"Budi", "Andi", "Eko", "Agus", "Rian", "Dedi", "Hendra", "Irwan", "Siti", "Dewi",
		"Rini", "Sari", "Mega", "Fitri", "Wulan", "Aditya", "Rizky", "Dimas", "Ahmad", "Taufik",
	}

	lastNames := []string{
		"Santoso", "Wijaya", "Saputra", "Pratama", "Hidayat", "Kusuma", "Setiawan", "Nugroho",
		"Lestari", "Utami", "Putri", "Rahmawati", "Permata", "Siregar", "Nasution", "Sitorus",
	}

	cities := []string{"Jakarta", "Bandung", "Surabaya", "Medan", "Semarang", "Yogyakarta", "Makassar", "Malang"}
	streets := []string{"Jl. Sudirman", "Jl. Gatot Subroto", "Jl. Merdeka", "Jl. Pemuda", "Jl. Diponegoro", "Jl. Ahmad Yani"}
	domains := []string{"gmail.com", "yahoo.com", "mail.id"}

	// 3. Loop untuk Insert 20 Data
	for i := 1; i <= 20; i++ {
		// Acak nama, alamat, dan email
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		fullName := fmt.Sprintf("%s %s", firstName, lastName)

		// Membuat email berdasarkan nama (lowercase tanpa spasi)
		emailBase := strings.ToLower(fmt.Sprintf("%s.%s%d", firstName, lastName, rand.Intn(99)))
		email := fmt.Sprintf("%s@%s", emailBase, domains[rand.Intn(len(domains))])

		passwordFake := "$2a$12$eImiTxAk4rfV..." // Simulasi hash bcrypt

		// Nomor telepon format Indonesia (08xx)
		phoneNumber := fmt.Sprintf("08%d%d", rand.Intn(900000)+100000, rand.Intn(900000)+100000)

		// Alamat format Indonesia
		address := fmt.Sprintf("%s No. %d, %s", streets[rand.Intn(len(streets))], rand.Intn(150)+1, cities[rand.Intn(len(cities))])

		// Memulai Transaksi Database
		tx, err := db.Begin()
		if err != nil {
			log.Printf("Gagal memulai transaksi ke-%d: %v\n", i, err)
			continue
		}

		// A. Insert ke tabel users
		var userID int
		queryUser := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
		err = tx.QueryRow(queryUser, email, passwordFake).Scan(&userID)
		if err != nil {
			tx.Rollback()
			log.Printf("[%d] Gagal insert user (%s): %v\n", i, email, err)
			continue
		}

		// B. Insert ke tabel user_profiles
		queryProfile := `INSERT INTO user_profiles (user_id, full_name, phone_number, address) 
						 VALUES ($1, $2, $3, $4)`
		_, err = tx.Exec(queryProfile, userID, fullName, phoneNumber, address)
		if err != nil {
			tx.Rollback()
			log.Printf("[%d] Gagal insert profile untuk User ID %d: %v\n", i, userID, err)
			continue
		}

		// Commit jika semua lancar
		err = tx.Commit()
		if err != nil {
			log.Printf("[%d] Gagal commit transaksi: %v\n", i, err)
		} else {
			fmt.Printf("[User %02d] ID: %-3d | %-25s | %-13s | %s\n", i, userID, email, phoneNumber, fullName)
		}
	}

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Proses pengisian 20 data users & profiles lokal Indonesia selesai!")
}
