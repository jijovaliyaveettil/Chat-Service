package handlers

import (
	"chat-service/initializers"
	"chat-service/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Find user by email first
	result := initializers.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
		"token":  tokenString,
	})
}

func Validate(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "This is the validate",
	})
}
