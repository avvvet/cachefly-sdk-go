package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Service ID argument
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Call GetBasic Service Options (GET /services/{id}/options)
	options, err := client.ServiceOptions.GetBasicOptions(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to get basic service options for %s: %v", serviceID, err)
	}

	// Pretty-print the options
	out, err := json.MarshalIndent(options, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting options JSON: %v", err)
	}

	fmt.Println("\n✅ Basic service options retrieved successfully:")
	fmt.Println(string(out))
}
