package handlers

import (
	"chat-service/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func CreateUser(ctx *gin.Context) {

	var req UserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Id:        uuid.NewString(),
		Name:      req.Name,
		UserName:  req.UserName,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	fmt.Println(user)

	ctx.JSON(200, gin.H{
		"status":   "ok",
		"username": req.UserName,
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
