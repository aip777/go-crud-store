package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
