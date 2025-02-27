package middleware

import (
	"chat-service/initializers"
	"chat-service/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Get the cookie of request
// Decode/validate it
// check the exp
// find the user exp
// attach to request
// continue
func AuthMiddleware(ctx *gin.Context) {
	// Get the token from the cookie
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	// Check for parsing errors
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Validate token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user by ID
		var user models.User
		result := initializers.DB.Where("id = ?", claims["sub"]).First(&user)
		if result.Error != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user to context
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
