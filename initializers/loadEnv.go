package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

}
