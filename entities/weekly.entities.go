package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Weekly struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
}

type WeeklyDeleteRequest struct {
	ID primitive.ObjectID `json:"id"`
}
