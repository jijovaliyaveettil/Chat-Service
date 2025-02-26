package initializers

import (
	"log"
	"os"
)

func LoadEnv() {

	// Explicitly get environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Validate required environment variables
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("Missing required database environment variables")
	}

}
