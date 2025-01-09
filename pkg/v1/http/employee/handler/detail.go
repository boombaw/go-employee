package handler

import (
	"go-employee/pkg/v1/constants"
	"go-employee/pkg/v1/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Detail struct {
	empService services.SrvEmployee
}

func (d *Detail) Handle(c *fiber.Ctx) (err error) {

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

	res, err := d.empService.FindByID(int64(empID))
	if err != nil {
		return
	}

	return c.Status(http.StatusOK).JSON(res)
}

func NewDetail(empService services.SrvEmployee) *Detail {
	return &Detail{
		empService: empService,
	}
}
