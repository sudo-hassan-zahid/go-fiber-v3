package database

import (
	"context"
	"fiber-v3/internal/config"
	"fiber-v3/internal/errors"
	"fiber-v3/internal/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.DBHost,
		cfg.DB.DBUser,
		cfg.DB.DBPassword,
		cfg.DB.DBName,
		cfg.DB.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Tokens{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func Health() error {
	if DB == nil {
		return errors.ErrDBNotInitialized
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return sqlDB.PingContext(ctx)
}
