package models

type Department struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	Code      string     `json:"code" gorm:"not null;uniqueIndex"`
	Employees []Employee `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
