package store

import (
	"github.com/masalennon/DIP_sample/db"
	"github.com/masalennon/DIP_sample/model"

)

func GetEmployeeByID(id string) (*model.Employee, error) {
	var e model.Employee
	if err := db.GetDB().Where(&model.Employee{ID: id}).First(&e).Error; err != nil {
		return nil, nil
	}

	return &e, nil
}