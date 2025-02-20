package endpoints

import (
	"chat-service/handlers"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(server *gin.Engine, version string) {
	a := server.Group("/user")
	a.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"version": version,
			"status":  "ok",
		})
	})
	a.POST("/create", handlers.CreateUser)
	a.GET("/get/:id", handlers.GetUser)
	a.PUT("/update/:id", handlers.UpdateUser)
	a.DELETE("/delete/:id", handlers.DeleteUser)
}
