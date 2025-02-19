package main

import (
	"fmt"

	"chat-service/endpoints"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0"

func main() {
	fmt.Println("Starting Server ........")

	server := gin.Default()

	endpoints.AddUserRoutes(server, version)

	server.Run(":8080")

}
