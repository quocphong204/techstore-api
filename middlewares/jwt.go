package middlewares

import (
	"net/http"
	"strings"
	"techstore-api/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 🔒 Check Redis blacklist
		blacklisted, err := config.RedisClient.Get(config.Ctx, "blacklist:"+tokenStr).Result()
		if err == nil && blacklisted == "true" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
			c.Abort()
			return
		}

		// ✅ Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Optional: extract claims nếu muốn
		// if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//     c.Set("user_id", claims["user_id"])
		// }

		c.Next()
	}
}
