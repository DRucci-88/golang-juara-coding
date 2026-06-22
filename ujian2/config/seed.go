package config

import (
	"gorm.io/gorm"

	"mini-hris/models"
)

func SeedData() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		departments := []models.Department{
			{Name: "Information Technology", Code: "DEPT-IT"},
			{Name: "Human Resources", Code: "DEPT-HR"},
			{Name: "Finance", Code: "DEPT-FIN"},
		}

		for _, department := range departments {
			if err := tx.Where(models.Department{Code: department.Code}).
				Assign(models.Department{Name: department.Name}).
				FirstOrCreate(&department).Error; err != nil {
				return err
			}
		}

		positions := []models.Position{
			{Title: "Software Engineer", BaseSalary: 8000000},
			{Title: "HR Specialist", BaseSalary: 6500000},
			{Title: "Finance Analyst", BaseSalary: 7000000},
		}

		for _, position := range positions {
			if err := tx.Where(models.Position{Title: position.Title}).
				Assign(models.Position{BaseSalary: position.BaseSalary}).
				FirstOrCreate(&position).Error; err != nil {
				return err
			}
		}

		var itDepartment models.Department
		var hrDepartment models.Department
		var engineerPosition models.Position
		var hrPosition models.Position

		if err := tx.Where("code = ?", "DEPT-IT").First(&itDepartment).Error; err != nil {
			return err
		}

		if err := tx.Where("code = ?", "DEPT-HR").First(&hrDepartment).Error; err != nil {
			return err
		}

		if err := tx.Where("title = ?", "Software Engineer").First(&engineerPosition).Error; err != nil {
			return err
		}

		if err := tx.Where("title = ?", "HR Specialist").First(&hrPosition).Error; err != nil {
			return err
		}

		employees := []models.Employee{
			{
				NIK:          "EMP-001",
				FullName:     "Andi Pratama",
				Email:        "andi.pratama@minihris.local",
				DepartmentID: itDepartment.ID,
				PositionID:   engineerPosition.ID,
				Status:       "ACTIVE",
				LeaveBalance: 12,
			},
			{
				NIK:          "EMP-002",
				FullName:     "Siti Lestari",
				Email:        "siti.lestari@minihris.local",
				DepartmentID: hrDepartment.ID,
				PositionID:   hrPosition.ID,
				Status:       "ACTIVE",
				LeaveBalance: 12,
			},
		}

		for _, employee := range employees {
			if err := tx.Where(models.Employee{NIK: employee.NIK}).
				Assign(models.Employee{
					FullName:     employee.FullName,
					Email:        employee.Email,
					DepartmentID: employee.DepartmentID,
					PositionID:   employee.PositionID,
					Status:       employee.Status,
					LeaveBalance: employee.LeaveBalance,
				}).
				FirstOrCreate(&employee).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
