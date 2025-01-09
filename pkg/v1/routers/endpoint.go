package routers

import (
	"go-employee/pkg/v1/domain/repository"
	department "go-employee/pkg/v1/http/department/handler"
	employee "go-employee/pkg/v1/http/employee/handler"
	"go-employee/pkg/v1/services"
	validators "go-employee/pkg/v1/validator"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Handle(c *fiber.Ctx) (err error)
}

var srvEmp = services.NewSrvEmployee(repository.NewEmployee())
var srvDept = services.NewSrvDepartment(repository.NewDepartment())
var validator = validators.NewXValidator()

var endpoint = map[string]Handler{
	"employee_list":          employee.NewList(srvEmp),
	"employee_detail":        employee.NewDetail(srvEmp),
	"employee_filter_status": employee.NewFilterStatus(srvEmp),
	"employee_create":        employee.NewCreate(srvEmp, validator),
	"employee_update":        employee.NewUpdate(srvEmp, validator),
	"salary":                 department.NewAverageSalary(srvDept),
}
