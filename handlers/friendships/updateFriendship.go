package handlers

import (
	"chat-service/initializers"
	"chat-service/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateFriendship(ctx *gin.Context) {
	user := getCurrentUserID(ctx)
	if user.Id == "" {
		return
	}

	// Validate request payload
	var request FriendshipRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate target friendship ID
	friendshipID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid friendship ID"})
		return
	}

	db := initializers.DB

	// Find the friendship record
	var friendship models.Friendships
	if err := db.First(&friendship, friendshipID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Friend request not found"})
		return
	}

	// Ensure that the logged-in user is the one receiving the request
	if friendship.FriendID != user.Id {
		ctx.JSON(403, gin.H{"error": "You are not authorized to respond to this request"})
		return
	}

	// Update the status
	friendship.Status = request.Status
	friendship.UpdatedAt = time.Now()

	if err := db.Save(&friendship).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update friendship status"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Friend request updated successfully", "status": request.Status})
}
