package routers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func APIKeyAuthMiddleware() fiber.Handler {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API Key is not set in environment variables")
	}

	return func(c *fiber.Ctx) error {
		clientAPIKey := c.Get("X-API-Key")
		if clientAPIKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API key required",
			})
		}
		if clientAPIKey != apiKey {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}
		return c.Next()
	}
}
