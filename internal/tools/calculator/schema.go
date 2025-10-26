package calculator

import (
	"fmt"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/muzfr7/mcp-toolkit/internal/tools/common"
)

// Operations constants.
const (
	OperationAdd      = "add"
	OperationSubtract = "subtract"
	OperationMultiply = "multiply"
	OperationDivide   = "divide"
)

// Params defines the input parameters for calculator operations.
type Params struct {
	X         float64 `json:"x" validate:"required"`
	Y         float64 `json:"y" validate:"required"`
	Operation string  `json:"operation" validate:"required,oneof=add subtract multiply divide"`
}

// Validate performs input validation.
func (p *Params) Validate() error {
	if p.Operation == "" {
		return common.ValidationError{
			Field:   "operation",
			Message: "operation is required",
		}
	}

	validOps := map[string]bool{
		OperationAdd:      true,
		OperationSubtract: true,
		OperationMultiply: true,
		OperationDivide:   true,
	}

	if !validOps[p.Operation] {
		return common.ValidationError{
			Field:   "operation",
			Message: fmt.Sprintf("invalid operation: %s", p.Operation),
		}
	}

	return nil
}

// GetSupportedOperations returns all supported operations.
func GetSupportedOperations() []string {
	return []string{
		OperationAdd,
		OperationSubtract,
		OperationMultiply,
		OperationDivide,
	}
}

// CreateSchema creates the JSON schema for calculator input.
func CreateSchema() *jsonschema.Schema {
	supportedOps := GetSupportedOperations()
	enumOps := make([]any, len(supportedOps))
	for i, op := range supportedOps {
		enumOps[i] = op
	}

	return &jsonschema.Schema{
		Type: "object",
		Properties: map[string]*jsonschema.Schema{
			"x": {
				Type:        "number",
				Description: "First number",
			},
			"y": {
				Type:        "number",
				Description: "Second number",
			},
			"operation": {
				Type:        "string",
				Description: "The arithmetic operation to perform",
				Enum:        enumOps,
			},
		},
		Required: []string{"x", "y", "operation"},
	}
}
