package controllers

import (
	"context"
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	item.CreatedAt = time.Now()

	_, err := config.DB.Collection("items").InsertOne(context.TODO(), item)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create item"})
		return
	}
	c.JSON(201, item)
}

func ListItems(c *gin.Context) {
	cursor, err := config.DB.Collection("items").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	var items []models.Item
	if err := cursor.All(context.TODO(), &items); err != nil {
		c.JSON(500, gin.H{"error": "Could not read items"})
		return
	}
	c.JSON(200, items)
}
