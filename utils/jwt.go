package utils

import (
	"errors"
	"time"
	"techstore-api/config" // ✅ import thêm

	"github.com/golang-jwt/jwt/v5"
)

// Không cần định nghĩa jwtKey nữa, dùng config.JWTSecret

func GenerateJWT(userID uint, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"name":    userName,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("phương thức ký không hợp lệ")
		}
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token không hợp lệ")
}
