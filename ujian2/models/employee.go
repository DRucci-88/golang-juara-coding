package models

type Employee struct {
	ID           uint         `json:"id" gorm:"primaryKey"`
	NIK          string       `json:"nik" gorm:"not null;uniqueIndex"`
	FullName     string       `json:"full_name" gorm:"not null"`
	Email        string       `json:"email" gorm:"not null;uniqueIndex"`
	DepartmentID uint         `json:"department_id" gorm:"not null"`
	PositionID   uint         `json:"position_id" gorm:"not null"`
	Status       string       `json:"status" gorm:"not null"`
	LeaveBalance int          `json:"leave_balance" gorm:"not null;default:12"`
	Department   Department   `json:"department,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Position     Position     `json:"position,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Attendances  []Attendance `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Leaves       []Leave      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Salaries     []Salary     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
