package handlers

import (
	"chat-service/initializers"
	"chat-service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendshipRequest struct {
	Status string `json:"status" binding:"required,oneof=pending accepted rejected"`
}

func getCurrentUserID(ctx *gin.Context) models.User {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return models.User{}
	}

	// Type assert to models.User and get the Id
	userModel, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
		return models.User{}
	}

	return userModel
}

func CreateFriendship(ctx *gin.Context) {
	user := getCurrentUserID(ctx)
	if user.Id == "" {
		return
	}

	// validate the target ID
	targetID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		fmt.Println("Invalid UUID:", err)
		ctx.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	// user should not be able to send a friend request to themselves
	if user.Id == targetID.String() {
		ctx.JSON(400, gin.H{"error": "Cannot send friend request to yourself"})
		return
	}

	var db = initializers.DB

	friendship := models.Friendships{
		UserID:   user.Id,
		FriendID: targetID.String(),
		Status:   "pending",
	}

	// create friendship
	if err := db.Create(&friendship).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Friend request failed"})
		return
	}

	ctx.JSON(201, friendship)
}
