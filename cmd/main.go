package main

import (
	"log"
	"os"

	"go1fl-sprint6/internal/server"
)

func main() {
	// Create logger
	logger := log.New(os.Stdout, "http-server: ", log.LstdFlags)

	// Create and start server
	srv := server.New(logger)

	logger.Printf("Starting server on :8080")
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
