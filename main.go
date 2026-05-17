package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "API Server",
	})

	app.Use(logger.New(logger.Config{
		Format:      "[${time}] ${status} - ${method} ${path}\n",
		ForceColors: true,
	}))

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}

}
