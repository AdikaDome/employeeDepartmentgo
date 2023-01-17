package database

import (
	"employee_department/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=postgres password=root dbname=book port=5432 sslmode=disable"

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		log.Println("connection successful")
	}
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(model.Employee{})

	return db

}
