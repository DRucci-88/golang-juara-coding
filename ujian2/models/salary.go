package models

type Salary struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	EmployeeID  uint     `json:"employee_id" gorm:"not null;uniqueIndex:idx_salary_employee_period"`
	Period      string   `json:"period" gorm:"not null;size:7;uniqueIndex:idx_salary_employee_period"`
	BasicSalary float64  `json:"basic_salary" gorm:"not null"`
	Allowance   float64  `json:"allowance" gorm:"not null"`
	Deductions  float64  `json:"deductions" gorm:"not null"`
	NetSalary   float64  `json:"net_salary" gorm:"not null"`
	Employee    Employee `json:"employee,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
