package main

import (
	"errors"
	"log"
	"net/http"
	"praktikum/app"
	"praktikum/middleware"
	"praktikum/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := app.NewDB()

	r := gin.Default()

	// Route Group dengan Proteksi Middleware
	v1 := r.Group("/api/v1")
	v1.Use(middleware.ApiKeyAuth())
	// r.Use(middleware.RateLimiter(5))

	v1.GET("/products", func(c *gin.Context) {
		var list []model.Product
		search := c.Query("search")

		if search != "" {
			db = db.Where("name ILIKE ?", "%s"+search+"%s")
		}

		if err := db.Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": list})
	})

	v1.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var p model.Product
		if err := db.Preload("Category").First(&p, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": p})
	})

	v1.POST("/products", func(c *gin.Context) {
		var req model.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&req).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product Success Created",
			"data":    req,
		})
	})

	v1.PUT("/products/:id", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))
		log.Println("Update Product By Id ", id)
		var p model.Product

		if err := db.First(&p, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Product not found",
				})
				return
			}
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		log.Printf("p %+v", p)

		var req model.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("req %+v", req)

		updateData := map[string]any{
			"category_id": req.CategoryID,
			"sku":         req.SKU,
			"name":        req.Name,
			"price":       req.Price,
			"stock":       req.Stock,
		}

		if err := db.Model(&p).Updates(updateData).Error; err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product Success Updated",
			"data":    p,
		})
	})

	v1.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var p model.Product
		if err := db.First(&p, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Melakukan Soft Delete
		if err := db.Delete(&p).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus secara lunak (Soft Deleted)"})
	})

	v1.GET("/categories/:id/products", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var category model.Category
		if err := db.First(&category, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category tidak ditemukan"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var list []model.Product

		if err := db.Preload("Category").Find(&list).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": list})
	})

	r.Run(":8080")
}
