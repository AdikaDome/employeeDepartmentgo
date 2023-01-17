package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"emp_name"`
	EmpSalary float32 `json:"emp_salary,omitempty"`
	EmpEmail  string  `json:"emp_email,omitempty"`
}
