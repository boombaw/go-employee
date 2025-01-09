package routers

import (
	"encoding/json"
	validators "go-employee/pkg/v1/validator"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/rs/zerolog/log"
)

// Route for mapping from json file
type Route struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Module   string `json:"module"`
	Endpoint string `json:"endpoint_filter"`
}

func Routes() {
	app := fiber.New(
		fiber.Config{
			// Global custom error handler
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusBadRequest).JSON(validators.GlobalErrorHandlerResp{
					Success: false,
					Message: err.Error(),
				})
			},
		},
	)

	app.Use(cors.New())
	app.Use(recover.New())
	// Gunakan middleware API Key Authentication
	app.Use(APIKeyAuthMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"title":   "Employe API v1",
			"message": "Welcome to employee managment v1",
		})
	})

	prefix := "/api"

	routers := loadRoutes("routes.json")
	for _, route := range routers {
		app.Add(route.Method, prefix+route.Path, endpoint[route.Endpoint].Handle)
	}

	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal().Err(err).Msg("Fiber app error")
	}
}

func loadRoutes(filePath string) []Route {
	var routes []Route

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal().
			Err(err).
			Msgf(" Failed to load file: %v", err)
	}

	if err := json.Unmarshal(file, &routes); err != nil {
		log.Fatal().Err(err).Msgf("Failed to marshal file : %v", err)
	}

	return routes
}
