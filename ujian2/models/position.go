package models

type Position struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Title      string     `json:"title" gorm:"not null"`
	BaseSalary float64    `json:"base_salary" gorm:"not null"`
	Employees  []Employee `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
