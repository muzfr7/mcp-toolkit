package mcpserver

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muzfr7/mcp-toolkit/internal/config"
	"github.com/muzfr7/mcp-toolkit/internal/logging"
	"github.com/muzfr7/mcp-toolkit/internal/tools/calculator"
)

// Run starts the MCP server with the provided configuration.
func Run(cfg *config.Config) error {
	// create MCP server instance
	srv := mcp.NewServer(&mcp.Implementation{
		Name:    cfg.Server.Name,
		Version: cfg.Server.Version,
	}, nil)

	// register tools
	if err := registerTools(srv); err != nil {
		return err
	}

	// create streamable HTTP handler
	mcpHandler := mcp.NewStreamableHTTPHandler(func(request *http.Request) *mcp.Server {
		return srv
	}, nil)

	// create HTTP mux
	mux := http.NewServeMux()

	// register Server-Sent Events endpoint handler
	mux.Handle("/sse", mcpHandler)

	// register health check endpoints handlers for Kubernetes
	mux.HandleFunc("/live", livenessHandler(cfg))
	mux.HandleFunc("/ready", readinessHandler(cfg))

	// wrap mux with HTTP logging middleware (now supports streaming)
	loggedHandler := logging.WithHTTPMiddleware(mux, slog.Default())

	slog.Info("Starting MCP server", "host", "localhost", "port", cfg.Server.Port)

	// start server
	addr := ":" + cfg.Server.Port
	if err := http.ListenAndServe(addr, loggedHandler); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server shutdown failed", "error", err)

		return err
	}

	slog.Info("Clean shutdown complete")

	return nil
}

// registerTools registers all available tools with the MCP server.
func registerTools(server *mcp.Server) error {
	calc := calculator.New()
	tool, handler := calc.Tool()

	// register calculator tool
	server.AddTool(tool, handler)

	return nil
}
