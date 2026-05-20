package handlers

import (
	"fiber-v3/internal/database"
	"time"

	"github.com/gofiber/fiber/v3"
)

func HealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "ok",
		"service":   "API Server",
		"timestamp": time.Now().UTC(),
	})
}

func DBHealthCheck(c fiber.Ctx) error {
	err := database.Health()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":    "error",
			"message":   "Database connection failed",
			"timestamp": time.Now().UTC(),
			"error":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "ok",
		"timestamp": time.Now().UTC(),
		"message":   "Database connection healthy",
	})
}
