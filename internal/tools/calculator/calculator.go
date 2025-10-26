package calculator

import (
	"errors"
	"fmt"
)

// Custom errors.
var (
	ErrDivisionByZero   = errors.New("division by zero is not allowed")
	ErrInvalidOperation = errors.New("invalid operation")
)

// Calculator provides arithmetic operations.
type Calculator struct{}

// New creates a new Calculator instance.
func New() *Calculator {
	return &Calculator{}
}

// Calculate performs the specified operation on two numbers.
func (c *Calculator) Calculate(x, y float64, operation string) (float64, error) {
	switch operation {
	case OperationAdd:
		return x + y, nil
	case OperationSubtract:
		return x - y, nil
	case OperationMultiply:
		return x * y, nil
	case OperationDivide:
		if y == 0 {
			return 0, ErrDivisionByZero
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrInvalidOperation, operation)
	}
}

// FormatResult formats the calculation result as a human-readable string.
func (c *Calculator) FormatResult(x, y, result float64, operation string) string {
	var symbol string
	switch operation {
	case OperationAdd:
		symbol = "+"
	case OperationSubtract:
		symbol = "-"
	case OperationMultiply:
		symbol = "×"
	case OperationDivide:
		symbol = "÷"
	default:
		symbol = "?"
	}

	return fmt.Sprintf("%.2f %s %.2f = %.2f", x, symbol, y, result)
}
