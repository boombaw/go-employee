package repository

import (
	"database/sql"
	"fmt"
	database "go-employee/pkg/v1/database/mysql"
	"go-employee/pkg/v1/database/query"
	"go-employee/pkg/v1/domain/models"

	"github.com/rs/zerolog/log"
)

type Employee interface {
	FindAll(page, limit int) (res []*models.Employee, err error)
	Create(arg models.EmployeeCreate) (res sql.Result, err error)
	FindByID(id int64) (res models.Employee, err error)
	CheckDuplicate(name, department string) (exist bool, err error)
	Update(arg models.Employee) (res sql.Result, err error)
	Paginate(page, limit int) (paginate interface{}, err error)
	FilterStatus(status string) (list []*models.Employee, err error)
}

type employee struct {
}

func NewEmployee() Employee {
	return &employee{}
}

func (e *employee) Create(arg models.EmployeeCreate) (res sql.Result, err error) {
	var db = database.DB

	tx := db.MustBegin()

	res = tx.MustExec(query.InsertEmployee, arg.Name, arg.Department, arg.Status, arg.Salary)

	log.Logger = log.With().Caller().Logger()

	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (e *employee) FindAll(page int, limit int) (list []*models.Employee, err error) {
	var db = database.DB

	var offset int

	if page > 1 {
		offset = ((page / 1) * limit) - limit
	}

	rows, err := db.Queryx(query.SelectAllEmp, limit, offset)
	if err != nil {
		log.Fatal().Msgf("Error while getting data : %v\n", err)
		return nil, err
	}

	for rows.Next() {
		var e models.Employee

		err = rows.StructScan(&e)
		if err != nil {
			return
		}

		list = append(list, &e)
	}

	return list, nil
}

func (e *employee) Paginate(page, limit int) (paginate interface{}, err error) {
	var db = database.DB
	var total int = 0

	err = db.Get(&total, query.TotalEmp)
	if err != nil {
		log.Fatal().Msg("Error while counting table")
		return nil, err
	}

	paginate = Paginate(total, page, limit)
	log.Info().Msgf("%v", paginate)
	return
}

func (e *employee) FindByID(id int64) (res models.Employee, err error) {
	var db = database.DB

	err = db.Get(&res, query.SelectEmpByID, id)
	if err != nil {
		return
	}

	return
}

func (e *employee) CheckDuplicate(name, department string) (exist bool, err error) {
	var count int
	var db = database.DB

	err = db.QueryRow(query.CheckDuplicateName, name, department).Scan(&count)
	if err != nil {
		return true, err
	}
	if count > 0 {
		err = fmt.Errorf("employee with name %s already exists in department %s", name, department)
		return true, err
	}

	return false, nil
}

func (e *employee) Update(arg models.Employee) (res sql.Result, err error) {
	var db = database.DB

	tx := db.MustBegin()

	res = tx.MustExec(query.UpdateEmployee, arg.Department, arg.Status, arg.Salary, arg.ID)

	log.Logger = log.With().Caller().Logger()

	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (e *employee) FilterStatus(status string) (list []*models.Employee, err error) {
	var db = database.DB

	log.Logger = log.With().Caller().Logger()

	rows, err := db.Queryx(query.SelectEmpByStatus, status)
	if err != nil {
		log.Fatal().Msgf("Error while getting data : %v\n", err)
		return nil, err
	}

	for rows.Next() {
		var e models.Employee

		err = rows.StructScan(&e)
		if err != nil {
			return
		}

		list = append(list, &e)
	}

	return list, nil
}
