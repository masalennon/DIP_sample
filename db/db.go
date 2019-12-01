package db

import (
	"fmt"

    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=test dbname=test password=postgres sslmode=disable")

	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}
