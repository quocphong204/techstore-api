package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"techstore-api/config"
)

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "🚫 No token provided"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// ✅ Giải mã token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "🚫 Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["exp"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "🚫 Invalid claims"})
		return
	}

	// ⏱️ Tính thời gian sống còn lại của token
	expUnix := int64(claims["exp"].(float64))
	ttl := time.Until(time.Unix(expUnix, 0))
	if ttl <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "⚠️ Token already expired"})
		return
	}

	// 🚫 Đưa token vào Redis blacklist
	err = config.RedisClient.Set(config.Ctx, "blacklist:"+tokenStr, "true", ttl).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Failed to blacklist token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ Logged out successfully"})
}
