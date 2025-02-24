package handlers

import (
	"chat-service/infra"
	"chat-service/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func CreateUser(ctx *gin.Context) {

	var req UserRequest

	db := infra.DB

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Id:        uuid.NewString(),
		Name:      req.Name,
		UserName:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	result := db.Create(&user)
	if result.Error != nil {
		// Handle specific errors
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(400, gin.H{"error": "Username or email already exists"})
			return
		}
		ctx.JSON(500, gin.H{"error": "Server error"})
		return
	}

	ctx.JSON(201, gin.H{
		"data": user,
	})
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	db := infra.DB

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

func UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	var req UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := infra.DB

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

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"status": "ok",
		"id":     userId,
	})
}
