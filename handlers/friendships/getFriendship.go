package handlers

import (
	"chat-service/initializers"
	"chat-service/models"

	"github.com/gin-gonic/gin"
)

func GetPendingFriendship(ctx *gin.Context) {
	user := getCurrentUserID(ctx)
	if user.Id == "" {
		return
	}

	var pendingRequests []models.Friendships
	db := initializers.DB

	// Fetch pending requests where the logged-in user is the FriendID (target)
	if err := db.Where("friend_id = ? AND status = ?", user.Id, "pending").Find(&pendingRequests).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve pending requests"})
		return
	}

	ctx.JSON(200, pendingRequests)
}
