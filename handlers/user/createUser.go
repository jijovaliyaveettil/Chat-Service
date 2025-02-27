package handlers

import (
	"chat-service/initializers"
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

	db := initializers.DB

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

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"status": "ok",
		"id":     userId,
	})
}
