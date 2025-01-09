package services

import (
	"database/sql"
	"go-employee/pkg/v1/constants"
	"go-employee/pkg/v1/domain/repository"
	"net/http"
)

type SrvDepartment interface {
	AvgSalary(department string) (res constants.DefaultResponse, err error)
}

type srvDepartment struct {
	empRepository repository.Department
}

func NewSrvDepartment(empRepository repository.Department) *srvDepartment {
	return &srvDepartment{
		empRepository: empRepository,
	}
}

func (a *srvDepartment) AvgSalary(department string) (res constants.DefaultResponse, err error) {

	avgSalary, err := a.empRepository.AvgSalary(department)
	if err != nil {
		if err == sql.ErrNoRows {
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: "No data found",
				Data:    nil,
				Errors:  nil,
			}
		} else {
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: constants.MESSAGE_FAILED,
				Data:    nil,
				Errors: []string{
					err.Error(),
				},
			}
		}
		err = nil
		return
	}
	res = constants.DefaultResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: constants.MESSAGE_SUCCESS,
		Data:    avgSalary,
	}

	return
}
