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

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	user.ID = primitive.NewObjectID()
	user.Token = "token-" + user.Username
	_, err := config.DB.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, user)
}

func LoginUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login"})
		return
	}

	var user models.User
	err := config.DB.Collection("users").FindOne(context.TODO(), bson.M{
		"username": input.Username,
		"password": input.Password,
	}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": user.Token, "user_id": user.ID})
}

func ListUsers(c *gin.Context) {
	cursor, err := config.DB.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		c.JSON(500, gin.H{"error": "Could not read users"})
		return
	}
	c.JSON(200, users)
}
