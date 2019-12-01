package employee

import (
	"github.com/masalennon/DIP_sample/model"

)

type Store interface {
	GetEmployeeByID(id string) (*model.Employee, error) 
}