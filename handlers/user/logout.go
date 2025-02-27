package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUser(ctx *gin.Context) {
	// Clear the Authorization cookie
	ctx.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Get token from Authorization header
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Blacklist the token
	// err := infra.BlacklistToken(tokenString, time.Hour*24)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Logged out successfully",
	})
}
