package endpoints

import "github.com/gin-gonic/gin"

func AddUserRoutes(server *gin.Engine, version string) {
	a := server.Group("/user")
	a.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"version": version,
			"status":  "ok",
		})
	})
}
