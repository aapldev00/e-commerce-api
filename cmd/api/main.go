// Package main is the entry point for the e-commerce API server.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aapldev00/e-commerce-api/platform/database"
	"github.com/joho/godotenv"
)

// main initializes the application infrastructure and starts the server.
func main() {
	// Initialize environment variables from .env file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Build the PostgreSQL connection string using environment variables.
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"localhost",
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Initialize the PostgreSQL connection pool.
	pool, err := database.NewPostgresPool(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize database pool: %v", err)
	}
	defer pool.Close()

	fmt.Println("🚀 Database connection established successfully")
}
