package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id"`
	Title     string               `json:"title" bson:"title"`
	Paragraph string               `json:"paragraph" bson:"paragraph"`
	Creator   primitive.ObjectID   `json:"creator" bson:"creator"`
	Comments  []primitive.ObjectID `json:"comments" bson:"comments"`
	Views     int32                `json:"views" bson:"views"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}

type BlogPostDeleteRequest struct {
	ID string `json:"_id"`
}
