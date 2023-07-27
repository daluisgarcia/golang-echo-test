package utils

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDB       string `env:"POSTGRES_DB,required"` // Allows to be read by .env
	PostgresHost     string `env:"POSTGRES_HOST,required"`
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	SecretKey        string `env:"SECRET_KEY,required"`
}

// Global configuration
var appConfig Config

func SetConfig() error {
	// Loading environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		return err
	}

	// Loading environment variables into the config struct
	err = env.Parse(&appConfig)

	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	// Returns the global configuration
	return appConfig
}