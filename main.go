package main

import (
	"fiber-v3/internal/config"
	"fiber-v3/internal/database"
	"fiber-v3/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "API Server",
	})

	app.Use(logger.New(logger.Config{
		Format:      "[${time}] ${status} - ${method} ${path}\n",
		ForceColors: true,
	}))

	routes.InitRoutes(app)

	app.Listen(":" + cfg.Port)

}
