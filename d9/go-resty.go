package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

// Definisikan struktur payload untuk request
type CategoryRequest struct {
	Name string `json:"name"`
}

type ProductRequest struct {
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

// Definisikan struktur response minimal untuk mengambil ID
type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductResponse struct {
	ID         uint `json:"id"`
	CategoryID uint `json:"category_id"`
}

func main() {
	// 1. Inisialisasi HTTP Client Resty
	client := resty.New().
		SetBaseURL("http://localhost:8080/api/v1"). // Ubah sesuai URL API Anda
		SetHeader("Content-Type", "application/json")

	// Data nama kategori dummy
	categories := []string{"Elektronik", "Pakaian", "Buku", "Kecantikan", "Olahraga"}

	fmt.Println("🚀 Memulai proses input data...")

	// 2. Loop untuk 5 Kategori
	for _, catName := range categories {
		var catResult CategoryResponse

		// Request membuat kategori
		resp, err := client.R().
			SetBody(CategoryRequest{Name: catName}).
			SetResult(&catResult). // Mapping langsung JSON ke struct
			Post("/categories")    // Ubah sesuai endpoint API Anda

		if err != nil {
			log.Fatalf("Gagal membuat kategori %s: %v", catName, err)
		}

		if resp.IsError() {
			log.Printf("⚠️ Server error saat membuat kategori %s: %s", catName, resp.String())
			continue
		}

		fmt.Printf("\n📂 Kategori Terbuat: %s (ID: %d)\n", catResult.Name, catResult.ID)

		// 3. Loop untuk 20 Produk di setiap Kategori
		for i := 1; i <= 20; i++ {
			var prodResult ProductResponse

			productPayload := ProductRequest{
				CategoryID:  catResult.ID,
				Name:        fmt.Sprintf("Produk %s %d", catResult.Name, i),
				Price:       float64(i * 15000), // Harga variasi, misal: 15.000, 30.000, dst.
				Stock:       10 + i,
				Description: fmt.Sprintf("Deskripsi lengkap untuk Produk %s nomor %d", catResult.Name, i),
			}

			// Request membuat produk
			pResp, pErr := client.R().
				SetBody(productPayload).
				SetResult(&prodResult).
				Post("/products") // Ubah sesuai endpoint API Anda

			if pErr != nil {
				log.Printf("❌ Gagal membuat produk %s ke-%d: %v", catResult.Name, i, pErr)
				continue
			}

			if pResp.IsError() {
				log.Printf("⚠️ Server error saat membuat produk ke-%d: %s", i, pResp.String())
				continue
			}

			fmt.Printf("   └─ 📦 Berhasil input: %s (ID Produk: %d)\n", productPayload.Name, prodResult.ID)
		}
	}

	fmt.Println("\n✅ Semua data kategori dan produk berhasil diproses!")
}
