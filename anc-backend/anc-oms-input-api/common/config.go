package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

// LoadChildEnv loads the environment variables from a specific file
func LoadChildEnv(filePath string) {
    if err := godotenv.Load(filePath); err != nil {
        log.Printf("No %s file found\n", filePath)
    }
}

// GetEnv gets an environment variable or returns a default value if not set
func GetEnv(key string, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}