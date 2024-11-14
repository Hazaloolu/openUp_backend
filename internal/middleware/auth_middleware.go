package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Ensure "Bearer " prefix is present and correctly sliced
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:] // Remove "Bearer " with space
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or Malformed token"})
			c.Abort()
			return
		}

		// Debug token string extraction
		fmt.Println("Extracted Token:", tokenString)

		// Validate the token
		claims, err := auth.ValidateJwt(tokenString)
		if err != nil || claims == nil {
			fmt.Println("Token validation error:", err) // Debugging line
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("UserID", claims.UserID)
		c.Next()
	}
}
