package handler

import (
	"go-employee/pkg/v1/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Status struct {
	empService services.SrvEmployee
}

func (d *Status) Handle(c *fiber.Ctx) (err error) {

	status := c.Query("status")
	if status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Missing required query parameters: status",
			},
		)
	}

	res, err := d.empService.FilterStatus(status)
	if err != nil {
		return
	}

	return c.Status(http.StatusOK).JSON(res)
}

func NewFilterStatus(empService services.SrvEmployee) *Status {
	return &Status{
		empService: empService,
	}
}
