package services

import (
	"database/sql"
	"fmt"
	"go-employee/pkg/v1/constants"
	"go-employee/pkg/v1/domain/models"
	"go-employee/pkg/v1/domain/repository"
	"net/http"

	"github.com/rs/zerolog/log"
)

type SrvEmployee interface {
	Create(req models.EmployeeCreate) (res constants.DefaultResponse, err error)
	Update(req models.Employee) (res constants.DefaultResponse, err error)
	FindAll(page, limit int) (res constants.DefaultResponse, err error)
	FindByID(id int64) (res constants.DefaultResponse, err error)
	FilterStatus(status string) (res constants.DefaultResponse, err error)
}

type srvEmployee struct {
	empRepository repository.Employee
}

func NewSrvEmployee(empRepository repository.Employee) *srvEmployee {
	return &srvEmployee{
		empRepository: empRepository,
	}
}

func (s *srvEmployee) Create(req models.EmployeeCreate) (res constants.DefaultResponse, err error) {
	var (
		result sql.Result
		lastID int64
		emp    models.Employee
	)

	duplicate, err := s.empRepository.CheckDuplicate(req.Name, req.Department)

	if !duplicate {
		newEmp := models.EmployeeCreate{
			Name:       req.Name,
			Department: req.Department,
			Status:     req.Status,
			Salary:     req.Salary,
		}

		result, err = s.empRepository.Create(newEmp)
		if err != nil {
			log.Err(err).Msg("Failed to create employee")
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: constants.MESSAGE_FAILED,
				Errors: []string{
					fmt.Sprintf("Failed to create employee: %v", err),
				},
			}
			return
		}
		lastID, err = result.LastInsertId()
		if err != nil {
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: constants.MESSAGE_FAILED,
				Errors:  []string{fmt.Sprintf("Failed to get last insert ID: %v", err)},
			}
			return
		}

		emp, err = s.empRepository.FindByID(lastID)

		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: constants.MESSAGE_SUCCESS,
			Data:    emp,
		}

	} else {
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: constants.MESSAGE_FAILED,
			Data:    nil,
			Errors: []string{
				fmt.Sprintf("%v", err),
			},
		}

		err = nil
	}

	return
}

func (s *srvEmployee) FindAll(page, limit int) (res constants.DefaultResponse, err error) {

	list, err := s.empRepository.FindAll(page, limit)
	if err != nil {
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: constants.MESSAGE_FAILED,
			Data:    nil,
			Errors: []string{
				err.Error(),
			},
		}
		err = nil
		return
	}

	paginate, _ := s.empRepository.Paginate(page, limit)
	res = constants.DefaultResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: constants.MESSAGE_SUCCESS,
		Data: map[string]interface{}{
			"employees":  list,
			"pagination": paginate,
		},
	}
	return
}

func (s *srvEmployee) FilterStatus(status string) (res constants.DefaultResponse, err error) {

	emp, err := s.empRepository.FilterStatus(status)
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
		Data:    emp,
	}

	return
}

func (s *srvEmployee) FindByID(id int64) (res constants.DefaultResponse, err error) {

	emp, err := s.checkExist(id)
	if err != nil {
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: constants.MESSAGE_FAILED,
			Data:    nil,
			Errors: []string{
				err.Error(),
			},
		}
		err = nil
		return
	}

	res = constants.DefaultResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: constants.MESSAGE_SUCCESS,
		Data:    emp,
	}

	return
}

func (s *srvEmployee) checkExist(id int64) (models.Employee, error) {
	var emp models.Employee

	emp, err := s.empRepository.FindByID(id)
	if err != nil {
		return emp, fmt.Errorf("employee with ID %d not found", id)
	}

	return emp, nil
}

func (s *srvEmployee) Update(req models.Employee) (res constants.DefaultResponse, err error) {

	_, err = s.checkExist(int64(req.ID))
	if err != nil {
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: constants.MESSAGE_FAILED,
			Data:    nil,
			Errors: []string{
				err.Error(),
			},
		}
		err = nil
		return
	}

	var duplicate bool
	duplicate, err = s.empRepository.CheckDuplicate(req.Name, req.Department)
	if err != nil {
		log.Err(err).Msg("Failed to check duplicate")
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: constants.MESSAGE_FAILED,
			Errors: []string{
				fmt.Sprintf("Failed to check duplicate: %v", err),
			},
		}
		return
	}

	if !duplicate {
		emp := models.Employee{
			ID:         req.ID,
			Name:       req.Name,
			Department: req.Department,
			Status:     req.Status,
			Salary:     req.Salary,
		}

		var result sql.Result
		result, err = s.empRepository.Update(emp)
		if err != nil {
			log.Err(err).Msg("Failed to upate employee")
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: constants.MESSAGE_FAILED,
				Errors: []string{
					constants.MESSAGE_FAILED,
				},
			}
			return
		}

		r, _ := result.RowsAffected()
		if r < 1 {
			// Handle case where no rows were affected
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: constants.MESSAGE_FAILED,
				Errors: []string{
					constants.MESSAGE_FAILED,
				},
			}
		} else {
			res = constants.DefaultResponse{
				Status:  http.StatusText(http.StatusOK),
				Message: constants.MESSAGE_SUCCESS,
				Data:    emp,
			}
		}

	} else {
		res = constants.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: constants.MESSAGE_FAILED,
			Data:    nil,
			Errors: []string{
				fmt.Sprintf("%v", err),
			},
		}
		err = nil
	}

	return
}
