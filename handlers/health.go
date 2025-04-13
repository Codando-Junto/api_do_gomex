// filepath: /Users/rgomes/code/pessoal/mentoria/api_do_gomex/handlers/health.go
package handlers

import "net/http"

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
