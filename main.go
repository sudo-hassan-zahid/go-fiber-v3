package main

import (
	"fiber-v3/internal/config"
	"fiber-v3/internal/routes"
	"log"

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

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	routes.InitRoutes(app)

	app.Listen(":" + cfg.Port)

}
