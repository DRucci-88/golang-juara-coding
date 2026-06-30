package model

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Title      string  `gorm:"type:varchar(50);not null"`
	BaseSalary float64 `gorm:"type:numeric(12,2);not null"`

	Employees []Employee `gorm:"foreignKey:PositionID"`
}

type PositionRepository[T any] interface {

	// INSERT INTO @@table
	Create(position *Position) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Position, error)

	// SELECT * FROM @@table
	FindAll() ([]Position, error)

	// UPDATE @@table
	Update(position *Position) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}
