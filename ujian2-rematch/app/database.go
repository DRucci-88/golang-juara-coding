package app

import (
	"ujian2_rematch/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "postgres://postgres:12345678@localhost:5432/ujian2_rematch?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Not Connected " + err.Error())
	}

	if err := db.AutoMigrate(
		&model.Attendance{},
		&model.Department{},
		&model.Employee{},
		&model.Leave{},
		&model.Position{},
		&model.Salary{},
	); err != nil {
		panic("Auto Migrate Failed " + err.Error())
	}

	return db
}
