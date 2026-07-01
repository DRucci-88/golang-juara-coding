package model

import (
	"errors"

	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Title      string  `gorm:"type:varchar(50);not null"`
	BaseSalary float64 `gorm:"type:numeric(12,2);not null"`

	Employees []Employee `gorm:"foreignKey:PositionID"`
}

type PositionPreload string

const (
	PositionPreloadEmployees PositionPreload = "Employees"
)

var (
	ErrPositionNotFound = errors.New("Position Not Found")
)
