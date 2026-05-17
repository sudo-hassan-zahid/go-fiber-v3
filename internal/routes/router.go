package routes

import "github.com/gofiber/fiber/v3"

func InitRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
