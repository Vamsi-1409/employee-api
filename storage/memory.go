package storage

import "employee-api/models"

var Employees = []models.Employee{
	{
		ID:         1,
		Name:       "Vamsi",
		Email:      "vamsi@example.com",
		Department: "DevOps",
	},
	{
		ID:         2,
		Name:       "Rahul",
		Email:      "rahul@example.com",
		Department: "Cloud",
	},
}
