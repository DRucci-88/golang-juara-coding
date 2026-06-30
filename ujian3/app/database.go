package app

import (
	"ujian3/helper"
	"ujian3/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "postgres://postgres:12345678@localhost:5432/ujian3?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Not Connected " + err.Error())
	}

	if err := db.AutoMigrate(
		&repository.UserDB{},
		&repository.EmployeeDB{},
		&repository.BlackListedTokenDB{},
		&repository.AttendanceDB{},
		&repository.DepartmentDB{},
		&repository.LeaveDB{},
		&repository.PositionDB{},
		&repository.SalaryDB{},
	); err != nil {
		panic("Auto Migrate Failed " + err.Error())
	}

	var count int64
	db.Model(&repository.UserDB{}).Count(&count)

	if count != 0 {
		return db
	}

	userDB := repository.UserDB{
		Email: "hrd1@company.co.id",
		Password: func() string {
			hash, _ := helper.HashPassword("hrd123456789")
			return hash
		}(),
		Role: "HRD",
	}
	db.Create(&userDB)

	// var count int64
	// db.Model(&model.User{}).Count(&count)

	// if count != 0 {
	// 	return db
	// }

	// user := model.User{
	// 	Email: "le.rucco@gmail.com",
	// 	Password: func() string {
	// 		hash, _ := helper.HashPassword("lerucco123456789")
	// 		return hash
	// 	}(),
	// }
	// db.Create(&user)

	// p1 := model.Product{
	// 	SKU:   "SKU-EL-1",
	// 	Name:  "keyboard Mechanical",
	// 	Price: 300_000,
	// 	Stock: 10,
	// }

	// p2 := model.Product{
	// 	SKU:   "SKU-EL-2",
	// 	Name:  "Mouse Wireless Gaming",
	// 	Price: 450_000,
	// 	Stock: 15,
	// }
	// db.Create(&p1)
	// db.Create(&p2)

	return db
}
