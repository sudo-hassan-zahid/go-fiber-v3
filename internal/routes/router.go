package routes

import (
	"fiber-v3/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	healthGroup := api.Group("/health")
	healthGroup.Get("/", handlers.HealthCheck)
	healthGroup.Get("/db", handlers.DBHealthCheck)

	v1 := api.Group("/v1")

	v1.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
