package handlers

import (
	"chat-service/infra"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendshipRequest struct {
	Status string `json:"status" binding:"required,oneof=pending accepted rejected"`
}

func CreateFriendship(ctx *gin.Context) {
	requesterID := getCurrentUserID(c)
	targetID := uuid.Parse(c.Param("id"))

	friendship := Friendship{
		UserID:   requesterID,
		FriendID: targetID,
		Status:   "pending",
	}

	if err := db.Create(&friendship).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Friend request failed"})
		return
	}

	ctx.JSON(201, friendship)
}

func UpdateFriendship(ctx *gin.Context) {
	userID := getCurrentUserID(c)
	requesterID := uuid.Parse(c.Param("id"))

	db := infra.DB

	// Find the pending request
	var friendship Friendship
	db.Where("user_id = ? AND friend_id = ? AND status = 'pending'",
		requesterID, userID).
		First(&friendship)

	// Update status
	db.Model(&friendship).Update("status", "accepted")

	// Create reciprocal relationship
	reciprocal := Friendship{
		UserID:   userID,
		FriendID: requesterID,
		Status:   "accepted",
	}
	db.Create(&reciprocal)

	ctx.JSON(200, gin.H{"status": "accepted"})
}

func GetFriendship(ctx *gin.Context) {

}

func DeleteFriendship(ctx *gin.Context) {
	userId := ctx.Param("id")
}
