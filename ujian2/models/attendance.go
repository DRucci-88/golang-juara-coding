package models

type Attendance struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	EmployeeID uint     `json:"employee_id" gorm:"not null;uniqueIndex:idx_attendance_employee_date"`
	Date       string   `json:"date" gorm:"not null;size:10;uniqueIndex:idx_attendance_employee_date"`
	CheckIn    string   `json:"check_in" gorm:"not null;size:5"`
	CheckOut   string   `json:"check_out" gorm:"not null;size:5"`
	Status     string   `json:"status" gorm:"not null"`
	Employee   Employee `json:"employee,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
