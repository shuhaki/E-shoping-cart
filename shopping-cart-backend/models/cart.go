package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Name      string             `bson:"name"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	Items     []CartItem         `bson:"items"`
}
