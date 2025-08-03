package main

import (
	"os"
	"shopping-cart-backend/config"
	"shopping-cart-backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.POST("/users/login", controllers.LoginUser)
		api.GET("/users", controllers.ListUsers)

		api.POST("/items", controllers.CreateItem)
		api.GET("/items", controllers.ListItems)

		api.POST("/carts", controllers.AddToCart)
		api.GET("/carts", controllers.GetCarts)

		api.POST("/orders", controllers.CreateOrder)
		api.GET("/orders", controllers.GetOrders)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
