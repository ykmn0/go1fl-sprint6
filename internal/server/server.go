package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ykmn0/go1fl-sprint6/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

// New creates a new HTTP server instance
func New(logger *log.Logger) *Server {
	// Create router and register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Create and configure the server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: srv,
	}
}

// Starts the HTTP server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
