package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. Struct Model Data & Validation
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"Price" binding:"required,gt=0"`
	Stock int     `json:"stock" binding:"required,min=0"`
}

var products = []Product{
	{ID: 1, Name: "Smartphone iPhone 15 Pro", Price: 21000000.00, Stock: 10},
	{ID: 2, Name: "TWS Sony WF-1000XM5", Price: 3500000.00, Stock: 25},
}
var nextID = 3

func main() {
	r := gin.Default()

	// Route Grouping
	api := r.Group("/api")
	{
		// GET /api/products - Mengambil semua produk dengan filter nama
		api.GET("/products", func(ctx *gin.Context) {
			search := ctx.Query("search")
			if search == "" {
				ctx.JSON(http.StatusOK, gin.H{"data": products})
				return
			}

			// Filter berdasarkan kata kunci nama
			var filtered []Product
			for _, p := range products {
				if containsIgnoreCase(p.Name, search) {
					filtered = append(filtered, p)
				}
			}
			ctx.JSON(http.StatusOK, gin.H{"data": filtered})
		})

		// GET /api/products/:id - Mengambil produk spesifik
		api.GET("/products/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk harus berupa angka"})
				return
			}
			for _, p := range products {
				if p.ID == id {
					c.JSON(http.StatusOK, gin.H{"data": p})
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		})

		// POST /api/products - Membuat produk baru
		api.POST("/products", func(c *gin.Context) {
			var newProduct Product
			if err := c.ShouldBindJSON(&newProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			newProduct.ID = nextID
			nextID++
			products = append(products, newProduct)
			c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil dibuat!", "data": newProduct})
		})

		// DELETE /api/products/:id - Menghapus produk
		api.DELETE("/products/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk harus berupa angka"})
				return
			}
			for idx, p := range products {
				if p.ID == id {
					// Hapus elemen dari slice products
					products = append(products[:idx], products[idx+1:]...)
					c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		})
	}

	r.Run(":8080")
}

// Fungsi helper sederhana untuk filter case-insensitive search
func containsIgnoreCase(src, search string) bool {
	// Konversi manual demi kesederhanaan helper
	return len(src) >= len(search) && (contains(stringsToLower(src),
		stringsToLower(search)))
}

func stringsToLower(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result = append(result, r+32)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
