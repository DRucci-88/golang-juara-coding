package models

type Leave struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	EmployeeID uint     `json:"employee_id" gorm:"not null"`
	StartDate  string   `json:"start_date" gorm:"not null;size:10"`
	EndDate    string   `json:"end_date" gorm:"not null;size:10"`
	Reason     string   `json:"reason" gorm:"not null"`
	Status     string   `json:"status" gorm:"not null"`
	Employee   Employee `json:"employee,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
