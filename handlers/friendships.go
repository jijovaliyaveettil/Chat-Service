package handlers

import (
	"chat-service/infra"
	"chat-service/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FriendshipRequest struct {
	Status string `json:"status" binding:"required,oneof=pending accepted rejected"`
}

func getCurrentUserID(ctx *gin.Context) string {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return ""
	}
	return userID.(string)
}

func CreateFriendship(ctx *gin.Context) {
	requesterID := getCurrentUserID(ctx)

	targetID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		fmt.Println("Invalid UUID:", err)
		ctx.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	var db = infra.DB

	friendship := models.Friendships{
		UserID:   requesterID,
		FriendID: targetID.String(),
		Status:   "pending",
	}

	if err := db.Create(&friendship).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Friend request failed"})
		return
	}

	ctx.JSON(201, friendship)
}

func UpdateFriendship(ctx *gin.Context) {
	userID := getCurrentUserID(ctx)
	requesterID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		fmt.Println("Invalid UUID:", err)
		ctx.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	db := infra.DB

	// Find the pending request
	var friendship models.Friendships
	db.Where("user_id = ? AND friend_id = ? AND status = 'pending'",
		requesterID, userID).
		First(&friendship)

	// Update status
	db.Model(&friendship).Update("status", "accepted")

	// Create reciprocal relationship
	reciprocal := models.Friendships{
		UserID:   userID,
		FriendID: requesterID.String(),
		Status:   "accepted",
	}
	db.Create(&reciprocal)

	ctx.JSON(200, gin.H{"status": "accepted"})
}

func GetFriendship(ctx *gin.Context) {

}

func DeleteFriendship(ctx *gin.Context) {
	// userId := ctx.Param("id")
}
