package controller

import (
	"employee_department/database"
	"employee_department/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func PostEmployee(response http.ResponseWriter, request *http.Request) {
	log.Printf("adding new employees .......................")

	response.Header().Set("Content-Type", "application/json")
	var Employee model.Employee
	json.NewDecoder(request.Body).Decode(&Employee)
	db := database.Connect()
	db.Create(&Employee)
	json.NewEncoder(response).Encode(Employee.ID)
	return

}

func GetEmployee(response http.ResponseWriter, request *http.Request) {
	log.Printf("get employee information")
	response.Header().Set("content-type", "application/json")
	db := database.Connect()
	params := mux.Vars(request)
	var Employee model.Employee
	db.First(&Employee, params["employeeId"])
	json.NewEncoder(response).Encode(Employee)
	return

}

func GetAllEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("all employees information")
	db := database.Connect()
	var employees []model.Employee
	db.Find(&employees)
	json.NewEncoder(response).Encode(employees)
	return

}

func UpdateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("update employee information")
	db := database.Connect()
	params := mux.Vars(request)
	var employee model.Employee
	db.First(&employee, params["employeeId"])
	json.NewDecoder(request.Body).Decode(&employee)
	db.Save(&employee)
	json.NewEncoder(response).Encode(employee)
	return

}

func DeleteAllEmployees(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("deleting all employees")
	db := database.Connect()
	var employees []model.Employee
	db.Find(&employees)
	db.Delete(&employees)
	json.NewEncoder(response).Encode("deleted all employees successfully")
	return
}

func DeleteEmployeeById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("deleting one employee by id")
	db := database.Connect()
	var Employee model.Employee
	params := mux.Vars(request)
	db.First(&Employee, params["employeeId"])
	db.Delete(&Employee)
	json.NewEncoder(response).Encode("deleted successfully")
	return
}
