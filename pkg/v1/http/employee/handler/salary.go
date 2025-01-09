package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Salary struct {
}

func (s *Salary) Handle(c *fiber.Ctx) (err error) {

	log.Info().Msg("Handle Salary employee")

	return c.JSON(fiber.Map{
		"data": "Salary Employe",
	})
}

func NewSalary() *Salary {
	return &Salary{}
}
