package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"techstore-api/config"
	"techstore-api/models"
)

// POST /products
func CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	stockStr := c.PostForm("stock")

	// ✅ Kiểm tra đầu vào rỗng
	if name == "" || priceStr == "" || stockStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Thiếu thông tin bắt buộc"})
		return
	}

	var price float64
	var stock int
	fmt.Sscanf(priceStr, "%f", &price)
	fmt.Sscanf(stockStr, "%d", &stock)

	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	// ✅ Xử lý ảnh nếu có
	file, err := c.FormFile("image")
	if err == nil {
		path := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err == nil {
			product.Image = "/" + path
		}
	}

	// ✅ Lưu vào DB
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Lưu DB thất bại: " + err.Error()})
		return
	}

	// Xóa cache Redis
	config.RedisClient.Del(config.Ctx, "products:all")

	c.JSON(http.StatusOK, gin.H{"message": "✅ Thêm sản phẩm thành công", "product": product})
}


// GET /products
func GetProducts(c *gin.Context) {
	cacheKey := "products:all"

	val, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var cached []models.Product
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			c.JSON(http.StatusOK, cached)
			return
		}
	}

	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Lỗi truy vấn DB: " + err.Error()})
		return
	}

	jsonData, _ := json.Marshal(products)
	config.RedisClient.Set(config.Ctx, cacheKey, jsonData, 5*time.Minute)

	c.JSON(http.StatusOK, products)
}
// PUT /products/:id
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "❌ Sản phẩm không tồn tại"})
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	stockStr := c.PostForm("stock")

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if priceStr != "" {
		fmt.Sscanf(priceStr, "%f", &product.Price)
	}
	if stockStr != "" {
		fmt.Sscanf(stockStr, "%d", &product.Stock)
	}

	// Cập nhật ảnh nếu có
	file, err := c.FormFile("image")
	if err == nil {
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		path := "uploads/" + fileName
		if err := c.SaveUploadedFile(file, path); err == nil {
			product.Image = "/" + path
		}
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Cập nhật thất bại"})
		return
	}

	// Xoá cache Redis
	config.RedisClient.Del(config.Ctx, "products:all")

	c.JSON(http.StatusOK, gin.H{"message": "✅ Cập nhật sản phẩm thành công", "product": product})
}
// DELETE /products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Xoá thất bại"})
		return
	}

	config.RedisClient.Del(config.Ctx, "products:all")

	c.JSON(http.StatusOK, gin.H{"message": "🗑️ Đã xoá sản phẩm"})
}
