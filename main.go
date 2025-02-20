package main

import (
	"fmt"
	"log"

	"chat-service/endpoints"
	"chat-service/infra"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var version = "1.0.0"

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	fmt.Println("Starting Server ........")

	// Initialize database connection
	infra.InitDatabase()

	server := gin.Default()

	endpoints.AddUserRoutes(server, version)

	// Start the server
	server.Run(":8080")
}
