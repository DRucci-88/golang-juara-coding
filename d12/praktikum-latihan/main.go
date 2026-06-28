package main

import (
	"praktikum/app"
	"praktikum/delivery"
	"praktikum/repository"
	"praktikum/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Daftarkan Custom SKU Validator
func registerSKUValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("sku_format", func(fl validator.FieldLevel) bool {
			sku, ok := fl.Field().Interface().(string)
			return ok && strings.HasPrefix(sku, "SKU-")
		})
	}
}
func main() {
	// Inisialisasi Database
	db := app.NewDatabase()

	// Auto Migration untuk model DB Repository
	db.AutoMigrate(&repository.ProductDB{})

	// Daftarkan validator kustom
	// registerSKUValidator()
	r := gin.Default()
	// 1. Dependency Injection (DI) - Merakit lapisan dari bawah ke atas
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	// 2. Inisialisasi HTTP Handler
	delivery.NewProductHandler(r, productUsecase)
	// Jalankan server
	r.Run(":8080")
}
 