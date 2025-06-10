package routes

import (
	"techstore-api/controllers"
	"techstore-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// ✅ Khai báo nhóm gốc
	api := r.Group("/api/v1")
	{
		api.POST("/login", controllers.Login)
		api.POST("/logout", controllers.Logout)
		api.POST("/register", controllers.Register)
		api.GET("/products", controllers.GetProducts)

		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			protected.POST("/products", controllers.CreateProduct)
			protected.PUT("/products/:id", controllers.UpdateProduct)
			protected.DELETE("/products/:id", controllers.DeleteProduct)

			protected.POST("/orders", controllers.CreateOrder)
			protected.GET("/categories", controllers.GetCategories)
			protected.POST("/categories", controllers.CreateCategory)

			protected.GET("/protected", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "✅ Authenticated successfully"})
			})

			admin := protected.Group("/admin")
			{
				admin.GET("/orders", controllers.GetAllOrders)
				admin.PUT("/orders/:id/status", controllers.UpdateOrderStatus)
			}
		}
	}
}


