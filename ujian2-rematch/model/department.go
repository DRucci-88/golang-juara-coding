package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
	Code string `gorm:"type:varchar(50);not null;uniqueIndex"`

	Employees []Employee `gorm:"foreignKey:DepartmentID"`
}

type DepartmentRepository[T any] interface {

	// INSERT INTO @@table
	Create(department *Department) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Department, error)

	// SELECT * FROM @@table
	FindAll() ([]Department, error)

	// SELECT * FROM @@table WHERE code = @code LIMIT 1
	FindByCode(code string) (*Department, error)

	// UPDATE @@table
	Update(department *Department) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}
