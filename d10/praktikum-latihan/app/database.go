package app

import (
	"praktikum/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// Data Source Name
	dsn := "postgres://postgres:12345678@localhost:5432/d10_praktikum?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal Koneksi dengan Database: " + err.Error())
	}

	// Jalankan Auto Migration
	err = db.AutoMigrate(&model.Product{})

	if err != nil {
		panic("Gagal Auto Migrate: " + err.Error())
	}

	return db
}
