package main

import (
	"log"
	"techstore-api/config"
	"techstore-api/middlewares"
	"techstore-api/models"
	"techstore-api/routes"
	"techstore-api/worker"
	"techstore-api/utils"

	_ "techstore-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ConnectDB()
	config.ConnectRedis()
	config.ConnectRabbitMQ()

	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Category{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
	// ✅ Tạo admin mặc định nếu chưa có
var count int64
	config.DB.Model(&models.User{}).Where("email = ?", "admin@gmail.com").Count(&count)
	if count == 0 {
	hashed, _ := utils.HashPassword("admin123")
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: hashed,
		Role:     "admin",
	}
	config.DB.Create(&admin)
	log.Println("✅ Admin mặc định đã được tạo: admin@gmail.com / admin123")
	}
	utils.SeedProducts()


	rabbitChan, err := config.RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open RabbitMQ channel: %v", err)
	}
	worker.StartOrderConsumer(rabbitChan)
	
	// ✅ Tạo router tại đây
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware()) // Middleware CORS để fix lỗi frontend gọi API
	r.Static("/uploads", "./uploads")
	
	// ✅ Truyền router vào setup routes
	routes.SetupRouter(r)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to run server: %v", err)
	}
}
