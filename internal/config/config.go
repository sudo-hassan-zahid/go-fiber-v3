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
	DB          DBConfig
}

var cfg *Config

func Init() {
	_ = godotenv.Load()

	cfg = &Config{
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENV"),
		DB: DBConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
		},
	}
}

func GetConfig() *Config {
	if cfg == nil {
		panic("config not initialized - call config.Init() first")
	}
	return cfg
}
