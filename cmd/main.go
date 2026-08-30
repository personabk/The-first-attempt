package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stderr, "INFO:", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Server started on :8080")

	if err := srv.Server.ListenAndServe(); err != nil {

		logger.Fatalf("Failed to start server: %v", err)
	}
}
