package utils

import (
	"fmt"
	"techstore-api/config"
	"techstore-api/models"
)

func SeedProducts() {
	var count int64
	config.DB.Model(&models.Product{}).Count(&count)
	if count > 0 {
		fmt.Println("🔁 Sản phẩm đã tồn tại, bỏ qua seed.")
		return
	}

	products := []models.Product{
		{
			Name:        "iPhone 15 Pro",
			Description: "Flagship mới nhất từ Apple",
			Price:       2699.99,
			Stock:       12,
			Image:       "/uploads/iphone.jpg", // ✅ Tên file ảnh đã có sẵn trong thư mục uploads/
		},
		{
			Name:        "Samsung Galaxy S23",
			Description: "Smartphone cao cấp của Samsung",
			Price:       1999.50,
			Stock:       15,
			Image:       "/uploads/samsung.jpg",
		},
		{
			Name:        "Xiaomi 13T",
			Description: "Giá tốt, cấu hình mạnh",
			Price:       799.99,
			Stock:       20,
			Image:       "/uploads/xiaomi.jpg",
		},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	fmt.Println("✅ Seed sản phẩm thành công.")
}
