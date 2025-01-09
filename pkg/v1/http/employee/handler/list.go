package handler

import (
	"go-employee/pkg/v1/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type List struct {
	empService services.SrvEmployee
}

func (l *List) Handle(c *fiber.Ctx) (err error) {

	log.Info().Msg("Handle list employee")

	limit := 10
	page, _ := strconv.Atoi(c.Params("page"))

	res, err := l.empService.FindAll(page, limit)
	if err != nil {
		return
	}

	return c.Status(http.StatusOK).JSON(res)
}

func NewList(empService services.SrvEmployee) *List {
	return &List{
		empService: empService,
	}
}
