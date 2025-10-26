package logging

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter wraps http.ResponseWriter to capture status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// WithHTTPMiddleware is a simple logging middleware based on the official MCP example.
func WithHTTPMiddleware(handler http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// generate request ID and add to context
		requestID := generateRequestID()
		ctx := WithRequestID(r.Context(), requestID)
		r = r.WithContext(ctx)

		// inject logger with request ID into context
		loggerWithID := logger.With(RequestIDField, requestID)
		ctx = WithLogger(ctx, loggerWithID)
		r = r.WithContext(ctx)

		// create response writer wrapper to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// add request ID to response headers for client debugging
		w.Header().Set("X-Request-ID", requestID)

		// log request details
		loggerWithID.Info("Request received",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"content_length", r.ContentLength)

		// call the actual handler
		handler.ServeHTTP(wrapped, r)

		// log response details
		duration := time.Since(start)
		logLevel := slog.LevelInfo
		if wrapped.statusCode >= 400 {
			logLevel = slog.LevelWarn
		}
		if wrapped.statusCode >= 500 {
			logLevel = slog.LevelError
		}

		loggerWithID.Log(r.Context(), logLevel, "Response sent",
			"method", r.Method,
			"path", r.URL.Path,
			"status_code", wrapped.statusCode,
			"duration_ms", duration.Milliseconds(),
			"duration", duration.String())
	})
}
