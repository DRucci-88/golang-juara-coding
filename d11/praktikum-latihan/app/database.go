package app

import (
	"praktikum/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "postgres://postgres:12345678@localhost:5432/d11_praktikum?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Not Connected " + err.Error())
	}

	if err := db.AutoMigrate(
		&model.Order{},
		&model.OrderItem{},
		&model.Product{},
		&model.User{},
	); err != nil {
		panic("Auto Migrate Failed " + err.Error())
	}

	var count int64
	db.Model(&model.User{}).Count(&count)

	if count != 0 {
		return db
	}

	userDummy := model.User{Email: "le.rucco@gmail.com"}
	db.Create(&userDummy)

	p1 := model.Product{
		SKU:   "SKU-EL-1",
		Name:  "keyboard Mechanical",
		Price: 300_000,
		Stock: 10,
	}

	p2 := model.Product{
		SKU:   "SKU-EL-2",
		Name:  "Mouse Wireless Gaming",
		Price: 450_000,
		Stock: 15,
	}
	db.Create(&p1)
	db.Create(&p2)

	return db
}
