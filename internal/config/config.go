package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type Config struct {
	Port        string
	Environment string
	DBConfig    DBConfig
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	cfg := &Config{
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
		DBConfig: DBConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
		},
	}

	return cfg, nil
}
