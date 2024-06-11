package config

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}