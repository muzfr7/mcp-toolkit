package calculator

import (
	"context"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muzfr7/mcp-toolkit/internal/logging"
)

// Tool returns the tool definition and handler with logger injection.
func (c *Calculator) Tool() (*mcp.Tool, mcp.ToolHandler) {
	// create tool definition
	definition := &mcp.Tool{
		Name:        "calculator",
		Description: "Performs basic arithmetic operations: add, subtract, multiply, divide",
		InputSchema: CreateSchema(),
	}

	// create base handler
	baseHandler := c.createHandler()

	// create wrapper handler that injects logger into context
	handler := func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// get the default logger (configured with server_name and environment)
		logger := slog.Default()

		// inject logger into context
		ctxWithLogger := logging.WithLogger(ctx, logger)

		// call the original handler with enhanced context
		return baseHandler(ctxWithLogger, request)
	}

	return definition, handler
}
