package model

import "gorm.io/gorm"

type Salary struct {
	gorm.Model

	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Period      string  `gorm:"type:varchar(10);not null"`
	BasicSalary float64 `gorm:"type:numeric(12,2);not null"`
	Allowance   float64 `gorm:"type:numeric(12,2);not null"`
	Deductions  float64 `gorm:"type:numeric(12,2);not null"`
	NetSalary   float64 `gorm:"type:numeric(12,2);not null"`
}

type SalaryRepository[T any] interface {

	// INSERT INTO @@table
	Create(salary *Salary) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Salary, error)

	// SELECT * FROM @@table
	FindAll() ([]Salary, error)

	// SELECT * FROM @@table WHERE employee_id = @employeeID
	FindByEmployeeID(employeeID uint) ([]Salary, error)

	// SELECT * FROM @@table WHERE period = @period
	FindByPeriod(period string) ([]Salary, error)

	// UPDATE @@table
	Update(salary *Salary) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}
