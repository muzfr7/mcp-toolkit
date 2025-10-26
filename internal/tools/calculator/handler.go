package calculator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muzfr7/mcp-toolkit/internal/logging"
)

// createHandler creates the MCP tool handler.
func (c *Calculator) createHandler() mcp.ToolHandler {
	return func(ctx context.Context, request *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// extract logger from context
		logger := logging.FromContext(ctx)

		// parse and validate input parameters
		var params Params
		if err := json.Unmarshal(request.Params.Arguments, &params); err != nil {
			return nil, fmt.Errorf("failed to parse arguments: %w", err)
		}

		if err := params.Validate(); err != nil {
			return nil, err
		}

		// perform calculation
		result, err := c.Calculate(params.X, params.Y, params.Operation)
		if err != nil {
			return nil, fmt.Errorf("calculation failed: %w", err)
		}

		// format result
		resultText := c.FormatResult(params.X, params.Y, result, params.Operation)

		// log calculation with structured data
		logger.Info("Calculator operation performed",
			"operation", params.Operation,
			"x", params.X,
			"y", params.Y,
			"result", result,
			"formatted_result", resultText)

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: resultText,
				},
			},
		}, nil
	}
}
