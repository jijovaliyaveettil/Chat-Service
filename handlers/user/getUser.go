package handlers

import (
	"chat-service/initializers"
	"chat-service/models"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	db := initializers.DB

	fetch_user := models.User{
		Id: userId,
	}
	result := db.First(&fetch_user)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   fetch_user,
	})
}
