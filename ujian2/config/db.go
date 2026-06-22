package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"mini-hris/models"
)

var DB *gorm.DB

func ConnectDatabase() error {
	db, err := gorm.Open(sqlite.Open("mini_hris.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return err
	}

	DB = db
	return autoMigrate()
}

func autoMigrate() error {
	return DB.AutoMigrate(
		&models.Department{},
		&models.Position{},
		&models.Employee{},
		&models.Attendance{},
		&models.Leave{},
		&models.Salary{},
	)
}
