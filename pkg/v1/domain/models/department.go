package models

// AvgSalary struct
type AvgSalary struct {
	Salary     float64 `json:"salary" db:"salary"`
	Department string  `json:"department" db:"department"`
}
