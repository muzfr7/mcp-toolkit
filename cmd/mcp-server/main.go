package main

import (
	"log/slog"
	"os"

	"github.com/muzfr7/mcp-toolkit/internal/config"
	"github.com/muzfr7/mcp-toolkit/internal/logging"
	"github.com/muzfr7/mcp-toolkit/internal/mcpserver"
)

func main() {
	// load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// setup logging
	logger := logging.NewLogger(logging.Config{
		ServerName:  cfg.Server.Name,
		Environment: cfg.Log.Environment,
		Level:       slog.LevelInfo, // default to Info for now
	})
	slog.SetDefault(logger)

	// start server
	if err := mcpserver.Run(cfg); err != nil {
		slog.Error("Server startup failed", "error", err)
		os.Exit(1)
	}
}
