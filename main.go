package main

import (
	"fmt"

	"chat-service/endpoints"
	"chat-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnv()
	initializers.InitDatabase()
}

var version = "1.0.0"

func main() {
	fmt.Println("Starting Server ........")

	// Create a new instance of a Gin router
	server := gin.Default()

	// Add user routes
	endpoints.AddUserRoutes(server, version)
	endpoints.AddFrienshipRoutes(server, version)

	// Start the server
	server.Run(":8080")
}
