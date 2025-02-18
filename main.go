package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	fmt.Println("Hello, World!")

	server := gin.Default()

	server.GET("/", getData)

	server.Run(":8080")

}
