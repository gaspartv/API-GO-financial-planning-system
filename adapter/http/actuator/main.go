package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody represents a health check
type HealthBody struct {
	Status string `json:"status"`
}

// Health returns a health check
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	status := HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
