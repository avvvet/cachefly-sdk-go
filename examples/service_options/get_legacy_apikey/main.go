// Example demonstrates retrieving the legacy API key for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Fetching legacy API key for a service
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id>
//
// Example:
//
//	go run main.go srv_123456789

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
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read service ID
	if len(os.Args) < 2 {
		log.Fatal("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the legacy API key
	key, err := client.ServiceOptions.GetLegacyAPIKey(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to get legacy API key for service %s: %v", serviceID, err)
	}

	// Print the key
	fmt.Println("✅ Legacy API key retrieved successfully:")
	jsonData, err := json.MarshalIndent(key, "", " ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [GetLegacyAPIKey]: %v", err)
	}
	fmt.Println(string(jsonData))
}
