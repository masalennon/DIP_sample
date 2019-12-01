package store

import (
	"github.com/jinzhu/gorm"
	"github.com/masalennon/DIP_sample/employee"
	"github.com/masalennon/DIP_sample/model"
)

type EmployeeGormStore struct {
	db *gorm.DB
}

func NewEmployeeGormStore(db *gorm.DB) employee.Store {
	return &EmployeeGormStore{
		db: db,
	}
}

func (es *EmployeeGormStore) GetEmployeeByID(id string) (*model.Employee, error) {
	var e model.Employee
	if err := es.db.Where(&model.Employee{ID: id}).First(&e).Error; err != nil {
		return nil, nil
	}

	return &e, nil
}
