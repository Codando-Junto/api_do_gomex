// filepath: /Users/rgomes/code/pessoal/mentoria/api_do_gomex/main.go
package main

import (
	"log"
	"net/http"

	"github.com/Codando-Junto/api_do_gomex/config"
	"github.com/Codando-Junto/api_do_gomex/handlers"
)

func main() {
	mux := http.NewServeMux() // Create a new ServeMux

	// Register routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	// Configure the server
	srv := config.NewServer(mux)

	// Log a message before starting the server
	log.Printf("Server is running on %s", srv.Addr)

	// Start the server
	log.Fatal(srv.ListenAndServe())
}
