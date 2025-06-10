package controllers

import (
	"net/http"
	"strconv"
	"techstore-api/config"
	"techstore-api/models"

	"github.com/gin-gonic/gin"
)

// Lấy toàn bộ đơn hàng (dành cho admin)
func GetAllOrders(c *gin.Context) {
	var orders []models.Order

	if err := config.DB.Preload("Items").Preload("User").Order("created_at DESC").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách đơn hàng"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Admin cập nhật trạng thái đơn hàng
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Đơn hàng không tồn tại"})
		return
	}

	order.Status = input.Status
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "✅ Cập nhật trạng thái đơn hàng thành công", "order": order})
}
