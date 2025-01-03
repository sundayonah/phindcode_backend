package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sundayonah/phindcode_backend/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

func NewConfig() (*ent.Client, error) {
	fmt.Println("Starting database connection")
	// Load environment variables
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		return nil, fmt.Errorf("database URL not found in environment variables")
	}

	// Open connection to PostgreSQL
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	// Set the global client
	Client = client

	return client, nil
}

func LoadEnv() error {
	return godotenv.Load()
}

func CloseDB() {
	Client.Close()
}
