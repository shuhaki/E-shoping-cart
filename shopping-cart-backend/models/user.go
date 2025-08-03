package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Token     string             `bson:"token"`
	CartID    *primitive.ObjectID `bson:"cart_id,omitempty"`
}
