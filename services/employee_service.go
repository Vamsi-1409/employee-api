package services

import (
	"employee-api/database"
	"employee-api/models"
)

func GetAllEmployees() []models.Employee {

	rows, err := database.DB.Query(
		"SELECT id, name, email, department FROM employees ORDER BY id",
	)
	if err != nil {
		return []models.Employee{}
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Department,
		)
		if err != nil {
			continue
		}

		employees = append(employees, employee)
	}

	return employees
}

func CreateEmployee(employee models.Employee) models.Employee {

	query := `
	INSERT INTO employees(name, email, department)
	VALUES($1, $2, $3)
	RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		employee.Name,
		employee.Email,
		employee.Department,
	).Scan(&employee.ID)

	if err != nil {
		return models.Employee{}
	}

	return employee
}

func GetEmployeeByID(id int) (*models.Employee, bool) {

	var employee models.Employee

	query := `
	SELECT id, name, email, department
	FROM employees
	WHERE id = $1
	`

	err := database.DB.QueryRow(query, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Department,
	)

	if err != nil {
		return nil, false
	}

	return &employee, true
}

func UpdateEmployee(id int, employee models.Employee) (*models.Employee, bool) {

	query := `
	UPDATE employees
	SET name = $1,
	    email = $2,
	    department = $3
	WHERE id = $4
	`

	result, err := database.DB.Exec(
		query,
		employee.Name,
		employee.Email,
		employee.Department,
		id,
	)

	if err != nil {
		return nil, false
	}

	rows, _ := result.RowsAffected()

	if rows == 0 {
		return nil, false
	}

	employee.ID = id

	return &employee, true
}

func DeleteEmployee(id int) bool {

	result, err := database.DB.Exec(
		"DELETE FROM employees WHERE id = $1",
		id,
	)

	if err != nil {
		return false
	}

	rows, _ := result.RowsAffected()

	return rows > 0
}
