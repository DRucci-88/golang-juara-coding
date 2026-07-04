package seeder

import (
	"database/sql"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"

	"ujian2_rematch/model"

	"gorm.io/gorm"
)

func seedAttendances(db *gorm.DB) error {

	var count int64

	if err := db.
		Model(&model.Attendance{}).
		Count(&count).
		Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("🌱 Attendance seed skipped")
		return nil
	}

	var employees []model.Employee

	if err := db.Find(&employees).Error; err != nil {
		return err
	}

	const (
		year  = 2026
		month = time.July
	)

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	var attendances []model.Attendance

	for _, employee := range employees {

		for date := start; date.Before(end); date = date.AddDate(0, 0, 1) {

			weekday := date.Weekday()

			attendance := model.Attendance{
				EmployeeID: employee.ID,
				Date:       date,
			}

			// Weekend
			if weekday == time.Saturday || weekday == time.Sunday {

				attendance.Status = model.AttendanceStatusAbsent

				attendances = append(attendances, attendance)

				continue
			}

			random := gofakeit.Number(1, 100)

			switch {

			// ==========================
			// PRESENT
			// ==========================
			case random <= 80:

				checkIn := time.Date(
					date.Year(),
					date.Month(),
					date.Day(),
					7,
					gofakeit.Number(50, 59),
					0,
					0,
					time.UTC,
				)

				checkOut := time.Date(
					date.Year(),
					date.Month(),
					date.Day(),
					17,
					gofakeit.Number(0, 30),
					0,
					0,
					time.UTC,
				)

				attendance.Status = model.AttendanceStatusPresent
				attendance.CheckIn = sql.NullTime{
					Time:  checkIn,
					Valid: true,
				}
				attendance.CheckOut = sql.NullTime{
					Time:  checkOut,
					Valid: true,
				}

			// ==========================
			// LATE
			// ==========================
			case random <= 95:

				checkIn := time.Date(
					date.Year(),
					date.Month(),
					date.Day(),
					8,
					gofakeit.Number(1, 45),
					0,
					0,
					time.UTC,
				)

				checkOut := time.Date(
					date.Year(),
					date.Month(),
					date.Day(),
					17,
					gofakeit.Number(0, 30),
					0,
					0,
					time.UTC,
				)

				attendance.Status = model.AttendanceStatusLate
				attendance.CheckIn = sql.NullTime{
					Time:  checkIn,
					Valid: true,
				}
				attendance.CheckOut = sql.NullTime{
					Time:  checkOut,
					Valid: true,
				}

			// ==========================
			// ABSENT
			// ==========================
			default:

				attendance.Status = model.AttendanceStatusAbsent
			}

			attendances = append(attendances, attendance)
		}
	}

	if err := db.CreateInBatches(&attendances, 500).Error; err != nil {
		return err
	}

	log.Printf("🌱 Seeded %d Attendances", len(attendances))

	return nil
}
