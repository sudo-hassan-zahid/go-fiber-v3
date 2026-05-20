package routes

import (
	"fiber-v3/internal/handlers"

	_ "fiber-v3/docs"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	healthGroup := api.Group("/health")
	healthGroup.Get("/", handlers.HealthCheck)
	healthGroup.Get("/db", handlers.DBHealthCheck)
	api.Get("/swagger/*", swaggo.HandlerDefault)
}
