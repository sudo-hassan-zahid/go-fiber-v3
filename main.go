package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "API Server",
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
