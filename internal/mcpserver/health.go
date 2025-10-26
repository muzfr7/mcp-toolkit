package mcpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/muzfr7/mcp-toolkit/internal/config"
	"github.com/muzfr7/mcp-toolkit/internal/logging"
)

// HealthStatus represents the health check response structure.
type HealthStatus struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

// livenessHandler creates a liveness check handler for liveness probes
// This endpoint checks if the application is running and not deadlocked.
func livenessHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.FromContext(r.Context())

		if r.Method != http.MethodGet {
			logger.Warn("Invalid method for liveness check", "method", r.Method)

			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := HealthStatus{
			Status:    "healthy",
			Timestamp: time.Now(),
			Service:   cfg.Server.Name,
			Version:   cfg.Server.Version,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("Failed to encode liveness response", "error", err)
		} else {
			logger.Debug("Liveness check successful")
		}
	}
}

// readinessHandler creates a readiness check handler for readiness probes
// This endpoint checks if the application is ready to serve traffic.
func readinessHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.FromContext(r.Context())

		if r.Method != http.MethodGet {
			logger.Warn("Invalid method for readiness check", "method", r.Method)

			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// TODO: Add actual readiness checks here, such as:
		// - Database connectivity
		// - External service dependencies
		// - Resource availability
		// - Tool initialization status
		// For now, we will assume the service is ready if it's running

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := HealthStatus{
			Status:    "ready",
			Timestamp: time.Now(),
			Service:   cfg.Server.Name,
			Version:   cfg.Server.Version,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("Failed to encode readiness response", "error", err)
		} else {
			logger.Debug("Readiness check successful")
		}
	}
}
