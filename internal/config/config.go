package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sundayonah/phindcode_backend/ent"

	_ "github.com/lib/pq"
)

type Config struct {
	Client *ent.Client
}

func NewConfig() *Config {
	fmt.Println("Starting database connection...")

	// Get database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Create a context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Open connection to database
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open connection to postgres: %v", err)
	}

	// Verify connection
	if err := client.Schema.Create(ctx); err != nil {
		client.Close()
		log.Fatalf("Failed to create schema resources: %v", err)
	}

	// Test the connection
	if err := client.Debug().Schema.Create(ctx); err != nil {
		client.Close()
		log.Fatalf("Failed to create schema: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	return &Config{
		Client: client,
	}
}

// CloseDB closes the database connection
func (c *Config) CloseDB() {
	if c.Client != nil {
		if err := c.Client.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}
