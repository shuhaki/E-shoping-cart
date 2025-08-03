package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
}
