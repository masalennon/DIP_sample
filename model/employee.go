package model

import (
)

type Employee struct {
	ID         string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Department string `gorm:"not null"`
}
