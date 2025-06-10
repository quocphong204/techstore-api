package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"techstore-api/config"
	"techstore-api/models"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ❌ Xóa cache khi có danh mục mới
	config.RedisClient.Del(config.Ctx, "categories:all")

	c.JSON(http.StatusOK, category)
}

func GetCategories(c *gin.Context) {
	cacheKey := "categories:all"

	// ✅ Kiểm tra cache
	val, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var cachedCategories []models.Category
		if err := json.Unmarshal([]byte(val), &cachedCategories); err == nil {
			c.JSON(http.StatusOK, cachedCategories)
			return
		}
	}

	// ❌ Nếu không có, truy vấn từ DB
	var categories []models.Category
	config.DB.Find(&categories)

	// 📦 Lưu vào cache
	data, _ := json.Marshal(categories)
	config.RedisClient.Set(config.Ctx, cacheKey, data, 5*time.Minute)

	c.JSON(http.StatusOK, categories)
}
