package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ItemID primitive.ObjectID `bson:"item_id"`
}
