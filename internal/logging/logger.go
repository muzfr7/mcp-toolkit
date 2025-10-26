package logging

import (
	"log/slog"
	"os"
)

const (
	// Field names for structured logging
	ServerNameField  = "server_name"
	EnvironmentField = "environment"
	RequestIDField   = "request_id"
)

// Config holds the logging configuration.
type Config struct {
	ServerName  string
	Environment string
	Level       slog.Level
}

// NewLogger creates a new structured JSON logger with the provided configuration.
func NewLogger(config Config) *slog.Logger {
	// set default level if not provided
	if config.Level == 0 {
		config.Level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: config.Level,
	})).With(
		ServerNameField, config.ServerName,
		EnvironmentField, config.Environment,
	)

	return logger
}
