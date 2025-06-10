package config

import "os"

var JWTSecret = getEnv("JWT_SECRET", "your-secret-key") // có thể đổi thành key bảo mật

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
