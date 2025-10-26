package logging

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type requestIDKey string

const requestIDContextKey requestIDKey = "request_id"

// generateRequestID generates a unique request ID.
func generateRequestID() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		// fallback to timestamp-based ID if crypto/rand fails
		return hex.EncodeToString(bytes)
	}

	return hex.EncodeToString(bytes)
}

// WithRequestID adds a request ID to the context.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDContextKey, requestID)
}

// RequestIDFromContext extracts the request ID from the context.
// Returns empty string if no request ID is found.
func RequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDContextKey).(string); ok {
		return requestID
	}

	return ""
}
