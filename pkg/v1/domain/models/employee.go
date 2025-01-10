package models

// Struct untuk Employee
type Employee struct {
	ID         int     `json:"id,omitempty" db:"id"`
	Name       string  `json:"name,omitempty"  db:"name"`
	Department string  `json:"department" validate:"required" db:"department"`
	Status     string  `json:"status" validate:"required" db:"status"`
	Salary     float64 `json:"salary" validate:"required" db:"salary"`
}

// Struct untuk Create Employee
type EmployeeCreate struct {
	Name       string  `json:"name" validate:"required,min=3" db:"name"`
	Department string  `json:"department" validate:"required" db:"department"`
	Status     string  `json:"status" validate:"required" db:"status"`
	Salary     float64 `json:"salary" validate:"required" db:"salary"`
}
