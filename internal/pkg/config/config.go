package config

import (
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds the application's configuration
type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig reads configuration from .env file and environment variables.
func LoadConfig() (*AppConfig, error) {
	// Load .env file, if it exist
	_ = godotenv.Load()

	return &AppConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}, nil
}
