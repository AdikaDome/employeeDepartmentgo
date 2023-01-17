package main

import (
	"employee_department/database"
	"employee_department/router"
)

func main() {
	database.Connect()
	router.Router()

}
