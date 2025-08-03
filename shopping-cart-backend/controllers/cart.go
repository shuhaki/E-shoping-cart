package controllers

import (
	"context"
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddToCartInput struct {
	UserID string `json:"user_id"`
	ItemID string `json:"item_id"`
}

func AddToCart(c *gin.Context) {
	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userId, _ := primitive.ObjectIDFromHex(input.UserID)
	itemId, _ := primitive.ObjectIDFromHex(input.ItemID)

	// Check if user already has a cart
	var cart models.Cart
	err := config.DB.Collection("carts").FindOne(context.TODO(), bson.M{
		"user_id": userId,
	}).Decode(&cart)

	// If no cart exists, create a new one
	if err != nil {
		cart = models.Cart{
			ID:        primitive.NewObjectID(),
			UserID:    userId,
			Name:      "My Cart",
			Status:    "open",
			CreatedAt: time.Now(),
			Items:     []models.CartItem{},
		}
	}

	cart.Items = append(cart.Items, models.CartItem{ItemID: itemId})

	// Upsert cart
	_, err = config.DB.Collection("carts").UpdateOne(context.TODO(),
		bson.M{"_id": cart.ID},
		bson.M{"$set": cart},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(200, cart)
}

func GetCarts(c *gin.Context) {
	cursor, err := config.DB.Collection("carts").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	var carts []models.Cart
	if err := cursor.All(context.TODO(), &carts); err != nil {
		c.JSON(500, gin.H{"error": "Could not read carts"})
		return
	}
	c.JSON(200, carts)
}
