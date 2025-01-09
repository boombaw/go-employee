package handler

import (
	"fmt"
	"go-employee/pkg/v1/constants"
	"go-employee/pkg/v1/domain/models"
	"go-employee/pkg/v1/services"
	validators "go-employee/pkg/v1/validator"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Update struct {
	empService services.SrvEmployee
	validator  validators.XValidator
}

func (u *Update) Handle(c *fiber.Ctx) (err error) {
	var req models.Employee

	reqID := c.Params("id")

	empID, err := strconv.Atoi(reqID)

	if err != nil {
		resp := constants.DefaultResponse{
			Status:  http.StatusText(http.StatusAccepted),
			Message: constants.MESSAGE_INVALID_REQUEST_FORMAT,
			Data:    nil,
			Errors: []string{
				"ID must be integer",
			},
		}
		return c.Status(http.StatusAccepted).JSON(resp)
	}

	req.ID = empID
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errs := u.validator.Validate(req)
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

	res, err := u.empService.Update(req)
	if err != nil {
		return
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func NewUpdate(empService services.SrvEmployee, validator *validators.XValidator) *Update {
	return &Update{
		empService: empService,
		validator:  *validator,
	}
}
