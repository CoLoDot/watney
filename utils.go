package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	if os.Getenv("ENV") == "PRODUCTION" || os.Getenv("ENV") == "STAGING" {
		return os.Getenv(key)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
