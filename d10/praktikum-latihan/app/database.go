package app

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"praktikum/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func randomPrices() float64 {
	return math.Round(float64(rand.Float64()*math.Pow(10, float64(rand.IntN(5)+2)))*100) / 100
}

func NewDB() *gorm.DB {
	// Data Source Name
	dsn := "postgres://postgres:12345678@localhost:5432/d10_praktikum?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal Koneksi dengan Database: " + err.Error())
	}

	// Jalankan Auto Migration
	err = db.AutoMigrate(&model.Product{}, &model.Category{})

	if err != nil {
		panic("Gagal Auto Migrate: " + err.Error())
	}

	var count int64
	db.Model(&model.Category{}).Count(&count)

	if count != 0 {
		return db
	}

	initialCategories := []model.Category{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
		{Name: "D"},
		{Name: "E"},
	}

	initialProducts := []model.Product{}

	for range 20 {
		initialProducts = append(initialProducts, model.Product{
			CategoryID: uint(rand.IntN(len(initialCategories)) + 1),
			SKU:        fmt.Sprintf("SKU-%d", rand.IntN(1000)+1),
			Name:       fmt.Sprintf("Name-%d", rand.IntN(1000)+1),
			Price:      randomPrices(),
			Stock:      rand.IntN(1000) + 1,
		})
	}

	errCategory := db.Create(&initialCategories).Error
	errProduct := db.Create(&initialProducts).Error

	if errCategory != nil || errProduct != nil {
		log.Println("Failed to seed initial")
		log.Println(errCategory)
		log.Println(errProduct)
		panic(map[string]any{
			"errProduct":  errProduct,
			"errCategory": errCategory,
		})
	}

	return db
}
