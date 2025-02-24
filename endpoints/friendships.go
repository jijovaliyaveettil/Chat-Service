package endpoints

import (
	"chat-service/handlers"

	"github.com/gin-gonic/gin"
)

func AddFrienshipRoutes(server *gin.Engine, version string) {
	a := server.Group("/friends")
	a.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"version": version,
			"status":  "ok",
		})
	})
	a.POST("/:id", handlers.CreateFriendship)
	a.PUT("/:id", handlers.UpdateFriendship)
	a.GET("/requests", handlers.GetFriendship)
	a.DELETE("/:id", handlers.DeleteFriendship)
}
