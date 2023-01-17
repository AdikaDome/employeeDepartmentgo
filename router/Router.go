package router

import (
	"employee_department/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", controller.PostEmployee).Methods("POST")
	router.HandleFunc("/employees", controller.GetAllEmployee).Methods("GET")
	router.HandleFunc("/employees/{employeeId}", controller.GetEmployee).Methods("GET")
	router.HandleFunc("/employees/{employeeId}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees", controller.DeleteAllEmployees).Methods("DELETE")
	router.HandleFunc("/employees/{employeeId}", controller.DeleteEmployeeById).Methods("DELETE")
	http.ListenAndServe(":8084", router)

}
