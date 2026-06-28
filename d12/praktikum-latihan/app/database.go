package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "postgres://postgres:12345678@localhost:5432/d12_praktikum?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database Not Connected " + err.Error())
	}

	if err := db.AutoMigrate(); err != nil {
		panic("Auto Migrate Failed " + err.Error())
	}

	return db
}
