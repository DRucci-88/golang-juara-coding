package repository

import (
	"time"
	"ujian3/domain"

	"gorm.io/gorm"
)

type EmployeeDB struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID uint `gorm:"uniqueIndex"`

	User *UserDB `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	NIK       string         `gorm:"type:varchar(16);unique;not null"`
	FullName  string         `gorm:"type:varchar(100);not null"`
	Status    string         `gorm:"type:varchar(20);default:'ACTIVE'"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (EmployeeDB) TableName() string { return "employees" }

func fromDomainEmployee(employee *domain.Employee) *EmployeeDB {
	return &EmployeeDB{
		ID:       employee.ID,
		UserID:   employee.UserID,
		NIK:      employee.NIK,
		FullName: employee.FullName,
		Status:   employee.Status,
	}
}

func toDomainEmployee(employeeDB *EmployeeDB) *domain.Employee {
	return &domain.Employee{
		ID:        employeeDB.ID,
		UserID:    employeeDB.UserID,
		NIK:       employeeDB.NIK,
		FullName:  employeeDB.FullName,
		Status:    employeeDB.Status,
		CreatedAt: employeeDB.CreatedAt,
		UpdatedAt: employeeDB.UpdatedAt,
		DeletedAt: employeeDB.DeletedAt.Time,
	}
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(
	db *gorm.DB,
) domain.EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r employeeRepository) Create(employee *domain.Employee) (*domain.Employee, error) {
	employeeDB := fromDomainEmployee(employee)

	if _, err := r.FindByNIK(employee.NIK); err != nil {
		return nil, domain.ErrEmployeeNIKDuplicate
	}

	err := r.db.Create(employeeDB).Error
	if err != nil {
		return nil, err
	}

	return toDomainEmployee(employeeDB), nil
}

func (r employeeRepository) FindByNIK(nik string) (*domain.Employee, error) {
	var employeeDB EmployeeDB
	err := r.db.Where("nik = ?", nik).First(&employeeDB).Error
	return toDomainEmployee(&employeeDB), err
}
