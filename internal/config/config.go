package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Server ServerConfig
	Log    LogConfig
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	Name    string `env:"SERVER_NAME,required"`
	Version string `env:"SERVER_VERSION,required"`
	Port    string `env:"PORT" envDefault:"8080"`
}

// LogConfig holds logging configuration.
type LogConfig struct {
	Level       string `env:"LOG_LEVEL,required"`
	Environment string `env:"ENV,required"`
}

// Load returns the application configuration.
func Load() (*Config, error) {
	// load .env file - return error if not found
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
