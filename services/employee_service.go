package services

import (
	"employee-api/models"
	"employee-api/storage"
)

func GetAllEmployees() []models.Employee {
	return storage.Employees
}

func CreateEmployee(employee models.Employee) models.Employee {

	employee.ID = len(storage.Employees) + 1

	storage.Employees = append(storage.Employees, employee)

	return employee
}

func GetEmployeeByID(id int) (*models.Employee, bool) {

	for _, employee := range storage.Employees {

		if employee.ID == id {
			return &employee, true
		}
	}

	return nil, false
}

func UpdateEmployee(id int, updatedEmployee models.Employee) (*models.Employee, bool) {

	for index, employee := range storage.Employees {

		if employee.ID == id {

			updatedEmployee.ID = id

			storage.Employees[index] = updatedEmployee

			return &storage.Employees[index], true
		}
	}

	return nil, false
}

func DeleteEmployee(id int) bool {

	for index, employee := range storage.Employees {

		if employee.ID == id {

			storage.Employees = append(
				storage.Employees[:index],
				storage.Employees[index+1:]...,
			)

			return true
		}
	}

	return false
}
