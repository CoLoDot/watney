package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable helper load variable regarding environnement type
func GetEnvVariable(key string) string {
	if os.Getenv("ENV") != "" {
		return os.Getenv(key)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
