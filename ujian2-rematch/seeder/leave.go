package seeder

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"

	"ujian2_rematch/model"

	"gorm.io/gorm"
)

func seedLeaves(db *gorm.DB) error {

	var count int64

	if err := db.
		Model(&model.Leave{}).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("🌱 Leave seed skipped")
		return nil
	}

	var employees []model.Employee

	if err := db.Find(&employees).Error; err != nil {
		return err
	}

	reasons := []string{
		"Family Vacation",
		"Medical Checkup",
		"Personal Leave",
		"Wedding Ceremony",
		"Family Event",
		"Child Graduation",
		"Religious Holiday",
		"Parents Visit",
	}

	var leaves []model.Leave

	for _, employee := range employees {

		// About 25% employees have leave
		if gofakeit.Number(1, totalEmployees) > 25 {
			continue
		}

		startDay := gofakeit.Number(2, 25)

		duration := gofakeit.Number(1, 3)

		startDate := time.Date(
			2026,
			time.July,
			startDay,
			0,
			0,
			0,
			0,
			time.UTC,
		)

		endDate := startDate.AddDate(0, 0, duration)

		status := gofakeit.RandomString([]string{
			string(model.LeaveStatusApproved),
			string(model.LeaveStatusApproved),
			string(model.LeaveStatusApproved),
			string(model.LeaveStatusPending),
			string(model.LeaveStatusRejected),
		})

		leave := model.Leave{
			EmployeeID: employee.ID,
			StartDate:  startDate,
			EndDate:    endDate,
			Reason:     reasons[gofakeit.Number(0, len(reasons)-1)],
			Status:     model.LeaveStatus(status),
		}

		leaves = append(leaves, leave)
	}

	if err := db.CreateInBatches(&leaves, 100).Error; err != nil {
		return err
	}

	log.Printf("🌱 Seeded %d Leaves", len(leaves))

	return nil
}
