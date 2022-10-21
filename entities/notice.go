package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
