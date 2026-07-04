package seeder

import (
	"log"
	"ujian2_rematch/model"

	"gorm.io/gorm"
)

func seedDepartments(db *gorm.DB) error {

	var count int64

	if err := db.
		Model(&model.Department{}).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("🌱 Department seed skipped")
		return nil
	}

	departments := []model.Department{
		{
			Name: "Information Technology",
			Code: "DEPT-IT",
		},
		{
			Name: "Human Resource",
			Code: "DEPT-HR",
		},
		{
			Name: "Finance",
			Code: "DEPT-FIN",
		},
		{
			Name: "Marketing",
			Code: "DEPT-MKT",
		},
		{
			Name: "Sales",
			Code: "DEPT-SLS",
		},
		{
			Name: "Procurement",
			Code: "DEPT-PRC",
		},
		{
			Name: "Operation",
			Code: "DEPT-OPS",
		},
		{
			Name: "Legal",
			Code: "DEPT-LGL",
		},
		{
			Name: "Customer Service",
			Code: "DEPT-CS",
		},
		{
			Name: "Research & Development",
			Code: "DEPT-RND",
		},
	}

	if err := db.CreateInBatches(&departments, 100).Error; err != nil {
		return err
	}

	log.Printf("🌱 Seeded %d Departments", len(departments))

	return nil
}
