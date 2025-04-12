package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loadEnv custom function to load env differently based on environment
func LoadEnv() {
	// Check if the app is running in a Docker environment
	environment := os.Getenv("APP_ENV")

	if environment == "staging" {
		// In staging/Docker, assume env variables are already set
		fmt.Println("Running in staging/Docker environment, loading .env file from /app")
		err := godotenv.Load("/app/.env")
		if err != nil {
			log.Fatalf("⚠️ Error loading .env file: %v", err)
		}
	} else {
		// In local development, load .env file
		fmt.Println("Running in development/local environment, loading .env file")
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("⚠️ Error loading .env file: %v", err)
		}
	}
}
