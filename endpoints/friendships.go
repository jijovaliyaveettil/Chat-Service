package endpoints

import (
	handlers "chat-service/handlers/friendships"
	"chat-service/middleware"

	"github.com/gin-gonic/gin"
)

func AddFrienshipRoutes(server *gin.Engine, version string) {
	a := server.Group("/friendships")
	a.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"version": version,
			"status":  "ok",
		})
	})
	a.POST("/:id", middleware.AuthMiddleware, handlers.CreateFriendship)
	a.PUT("/:id", middleware.AuthMiddleware, handlers.UpdateFriendship)
	a.GET("/requests", middleware.AuthMiddleware, handlers.GetFriendship)
	// a.DELETE("/:id", handlers.DeleteFriendship)
}
