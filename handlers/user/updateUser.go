package handlers

import (
	"chat-service/initializers"
	"chat-service/models"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	var req UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := initializers.DB

	update_user := models.User{
		Id: userId,
	}
	result := db.Model(&update_user).Updates(&req)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   update_user,
	})
}
