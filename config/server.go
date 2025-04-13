package config

import (
	"net/http"
	"time"
)

func NewServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
}
