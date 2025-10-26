package logging

import (
	"context"
	"log/slog"
)

type contextKey string

const loggerContextKey contextKey = "logger"

// WithLogger adds a logger to the context.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// FromContext extracts the logger from the context
// If no logger is found, it returns the default logger.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		return logger
	}

	return slog.Default()
}
