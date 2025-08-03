package controllers

import (
	"context"
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderInput struct {
	UserID string `json:"user_id"`
}

func CreateOrder(c *gin.Context) {
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userId, _ := primitive.ObjectIDFromHex(input.UserID)

	// Find user's cart
	var cart models.Cart
	err := config.DB.Collection("carts").FindOne(context.TODO(), bson.M{
		"user_id": userId,
		"status":  "open",
	}).Decode(&cart)

	if err != nil {
		c.JSON(404, gin.H{"error": "No open cart found"})
		return
	}

	// Create order
	order := models.Order{
		ID:        primitive.NewObjectID(),
		CartID:    cart.ID,
		UserID:    userId,
		CreatedAt: time.Now(),
	}

	_, err = config.DB.Collection("orders").InsertOne(context.TODO(), order)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	// Mark cart as closed
	_, _ = config.DB.Collection("carts").UpdateOne(context.TODO(), bson.M{
		"_id": cart.ID,
	}, bson.M{
		"$set": bson.M{"status": "ordered"},
	})

	c.JSON(200, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

func GetOrders(c *gin.Context) {
	cursor, err := config.DB.Collection("orders").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	var orders []models.Order
	if err := cursor.All(context.TODO(), &orders); err != nil {
		c.JSON(500, gin.H{"error": "Could not read orders"})
		return
	}
	c.JSON(200, orders)
}
