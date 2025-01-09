package handler

import (
	"fmt"
	"go-employee/pkg/v1/domain/models"
	"go-employee/pkg/v1/services"
	validators "go-employee/pkg/v1/validator"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Create struct {
	empService services.SrvEmployee
	validator  validators.XValidator
}

func (cr *Create) Handle(c *fiber.Ctx) (err error) {

	var req models.EmployeeCreate
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errs := cr.validator.Validate(req)
	if len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Message,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " dan "),
		}
	}

	res, err := cr.empService.Create(req)
	if err != nil {
		return
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func NewCreate(empService services.SrvEmployee, validator *validators.XValidator) *Create {
	return &Create{
		empService: empService,
		validator:  *validator,
	}
}
