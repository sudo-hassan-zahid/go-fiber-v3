package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fiber-v3/internal/config"
	"fiber-v3/internal/database"
	"fiber-v3/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// @title 				Fiber API
// @version 			1.0
// @description 		This is a sample API
func main() {
	config.Init()
	cfg := config.GetConfig()

	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "API Server - Mode: " + cfg.Environment,
	})

	app.Use(logger.New(logger.Config{
		Format:      "[${time}] ${status} - ${method} ${path}\n",
		ForceColors: true,
	}))

	routes.InitRoutes(app)

	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()

	log.Printf("Server running on port %s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server shut down cleanly")
}
