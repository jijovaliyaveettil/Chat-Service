package handlers

import (
	"chat-service/initializers"
	"chat-service/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateFriendship(ctx *gin.Context) {
	user := getCurrentUserID(ctx)
	if user.Id == "" {
		return
	}

	requesterID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		fmt.Println("Invalid UUID:", err)
		ctx.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	db := initializers.DB

	// Find the pending request
	var friendship models.Friendships
	db.Where("user_id = ? AND friend_id = ? AND status = 'pending'",
		requesterID, user.Id).
		First(&friendship)

	if friendship.ID == 0 {
		ctx.JSON(404, gin.H{"error": "Friend request not found"})
		return
	}
	// Update status
	db.Model(&friendship).Update("status", "accepted")

	// Create reciprocal relationship
	reciprocal := models.Friendships{
		UserID:   user.Id,
		FriendID: requesterID.String(),
		Status:   "accepted",
	}
	db.Create(&reciprocal)

	ctx.JSON(200, gin.H{"status": "accepted"})
}
