package seeder

import (
	"log"

	"ujian2_rematch/model"

	"gorm.io/gorm"
)

func seedPositions(db *gorm.DB) error {

	var count int64

	if err := db.
		Model(&model.Position{}).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("🌱 Position seed skipped")
		return nil
	}

	positions := []model.Position{
		{
			Title:      "Junior Software Engineer",
			BaseSalary: 6_500_000,
		},
		{
			Title:      "Software Engineer",
			BaseSalary: 8_500_000,
		},
		{
			Title:      "Senior Software Engineer",
			BaseSalary: 12_000_000,
		},
		{
			Title:      "Tech Lead",
			BaseSalary: 18_000_000,
		},
		{
			Title:      "HR Staff",
			BaseSalary: 6_000_000,
		},
		{
			Title:      "HR Specialist",
			BaseSalary: 8_000_000,
		},
		{
			Title:      "Finance Staff",
			BaseSalary: 6_500_000,
		},
		{
			Title:      "Finance Supervisor",
			BaseSalary: 10_000_000,
		},
		{
			Title:      "Marketing Executive",
			BaseSalary: 7_500_000,
		},
		{
			Title:      "Sales Executive",
			BaseSalary: 7_000_000,
		},
		{
			Title:      "Sales Manager",
			BaseSalary: 15_000_000,
		},
		{
			Title:      "Procurement Officer",
			BaseSalary: 7_000_000,
		},
		{
			Title:      "Operation Staff",
			BaseSalary: 6_500_000,
		},
		{
			Title:      "Customer Service",
			BaseSalary: 5_500_000,
		},
		{
			Title:      "Legal Officer",
			BaseSalary: 9_000_000,
		},
	}

	if err := db.CreateInBatches(&positions, 100).Error; err != nil {
		return err
	}

	log.Printf("🌱 Seeded %d Positions", len(positions))

	return nil
}
