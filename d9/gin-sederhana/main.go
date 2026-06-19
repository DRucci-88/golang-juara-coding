package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`  // Harus diisi dan > 0
	Stock int     `json:"stock" binding:"required,min=0"` // Harus diisi dan >= 0
}

func main() {
	// Inisialisasi engine default Gin (termasuk Logger & Recovery Middleware bawaan)
	r := gin.Default()

	// Routing sederhana
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	// Pengelompokan Route untuk Versi API
	v1 := r.Group("/api/v1")
	{
		// HTTP GET: Membaca data
		v1.GET("/products", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"data": "Semua product"})
		})

		// HTTP POST: Membuat data baru
		v1.POST("/products", func(ctx *gin.Context) {
			ctx.JSON(201, gin.H{"message": "Produk berhasil dibuat"})
		})

		// HTTP PUT: Mengubah Data Secara Keseluruhan
		v1.PUT("/products", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Produk berhasil diperbarui"})
		})
		// HTTP DELETE: Menghapus Data
		v1.DELETE("/products", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Produk berhasil dihapus"})
		})
	}

	/// 1. Mengambil Query Parameters (ctx.Query)
	r.GET("/search", func(ctx *gin.Context) {
		// http://localhost:8080/search?keyword=sepatu&page=2
		keyword := ctx.Query("keyword")
		page := ctx.DefaultQuery("page", "1")

		ctx.JSON(200, gin.H{
			"keyword": keyword,
			"page":    page,
		})
	})

	/// 2. Mengambil Path/URL Parameters ( ctx.Param )
	r.GET("/products/:id", func(ctx *gin.Context) {
		// http://localhost:8080/api/v1/products/45
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"product_id": id,
		})
	})

	r.POST("/products", func(ctx *gin.Context) {
		var req ProductRequest

		ccc := ctx.BindJSON(&req)
		fmt.Println(ccc)

		// Validasi dan Binding JSON ke struc ProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(), // Mengembalikan pesan kesalahn
			})
			return
		}

		// Jika validasi sukses, lakukan sesuatu dengan data req
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Produk valid!",
			"data":    req,
		})
	})

	/// 3. Binding Request Body JSON ( c.ShouldBindJSON )

	// Jalankan server di 8080
	r.Run(":8080")
}
