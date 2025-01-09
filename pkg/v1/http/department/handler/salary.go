package handler

import (
	"go-employee/pkg/v1/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type AverageSalary struct {
	deptService services.SrvDepartment
}

func (l *AverageSalary) Handle(c *fiber.Ctx) (err error) {

	log.Info().Msg("Handle Average Salary Department")

	department := c.Query("name")
	if department == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Missing required query parameters: name",
			},
		)
	}

	res, err := l.deptService.AvgSalary(department)
	if err != nil {
		return
	}

	return c.Status(http.StatusOK).JSON(res)
}

func NewAverageSalary(deptService services.SrvDepartment) *AverageSalary {
	return &AverageSalary{
		deptService: deptService,
	}
}
