package config

import (
    "github.com/joho/godotenv"
    "log"
)

// Function to load the config on the .env file (it is called on the main.go)
func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}
