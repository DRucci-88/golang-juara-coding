package seeder

import (
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"

	"ujian2_rematch/helper"
	"ujian2_rematch/model"

	"gorm.io/gorm"
)

const totalEmployees = 120

func seedUsersAndEmployees(db *gorm.DB) error {

	var count int64

	if err := db.
		Model(&model.Employee{}).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("🌱 Employee seed skipped")
		return nil
	}

	gofakeit.Seed(88)

	var departments []model.Department
	if err := db.Find(&departments).Error; err != nil {
		return err
	}

	var positions []model.Position
	if err := db.Find(&positions).Error; err != nil {
		return err
	}

	if len(departments) == 0 {
		return fmt.Errorf("department seed not found")
	}

	if len(positions) == 0 {
		return fmt.Errorf("position seed not found")
	}

	password, err := helper.HashPassword("password")
	if err != nil {
		return err
	}

	// =====================================================
	// ADMIN
	// =====================================================

	admin := model.User{
		Email:    "admin@company.co.id",
		Password: password,
		Role:     model.UserRoleAdmin,
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	// =====================================================
	// EMPLOYEES
	// =====================================================

	for i := 1; i <= totalEmployees; i++ {

		user := model.User{
			Email:    fmt.Sprintf("employee%03d@company.co.id", i),
			Password: password,
			Role:     model.UserRoleEmployee,
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}

		department := departments[gofakeit.Number(0, len(departments)-1)]
		position := positions[gofakeit.Number(0, len(positions)-1)]

		employee := model.Employee{
			UserID: user.ID,

			NIK: fmt.Sprintf(
				"EMP%05d",
				i,
			),

			FullName: gofakeit.Name(),

			Status: model.EmployeeStatusActive,

			DepartmentID: department.ID,
			PositionID:   position.ID,
		}

		if err := db.Create(&employee).Error; err != nil {
			return err
		}
	}

	log.Printf("🌱 Seeded %d Employees", totalEmployees)
	log.Printf("🌱 Admin Email    : admin@company.co.id")
	log.Printf("🌱 Admin Password : password")

	return nil
}
