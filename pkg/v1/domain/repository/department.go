package repository

import (
	database "go-employee/pkg/v1/database/mysql"
	"go-employee/pkg/v1/database/query"
	"go-employee/pkg/v1/domain/models"

	"github.com/rs/zerolog/log"
)

type Department interface {
	AvgSalary(department string) (res models.AvgSalary, err error)
}

type department struct {
}

func NewDepartment() Department {
	return &department{}
}

func (d *department) AvgSalary(department string) (res models.AvgSalary, err error) {
	var db = database.DB
	log.Logger = log.With().Caller().Logger()

	err = db.Get(&res, query.AvgSalary, department)
	log.Err(err).Msgf("Average Salary")
	if err != nil {
		return
	}
	return res, nil
}
