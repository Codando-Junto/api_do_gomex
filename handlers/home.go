package handlers

import (
	"encoding/json"
	"net/http"
)

// HomeResponse represents the structure of the JSON response
type Response struct {
	Message string `json:"message"`
}

// HomeHandler handles requests to the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Create the response data
	response := Response{
		Message: "Amo vocês!",
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Serialize the response to JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}
