package handlers

import "github.com/gin-gonic/gin"

func CreateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"status": "ok",
		"id":     userId,
	})
}

func UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"status": "ok",
		"id":     userId,
	})
}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"status": "ok",
		"id":     userId,
	})
}
