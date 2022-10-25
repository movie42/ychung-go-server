package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Groups struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string
	IsPublic  bool
	Groups    []primitive.ObjectID `json:"groups" bson:"groups"`
	Creator   primitive.ObjectID   `json:"creator" bson:"creator"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Group struct {
	ID        primitive.ObjectID `json:"_id" bson:"id"`
	Name      string
	Place     string
	Type      string
	HumanIds  []primitive.ObjectID `json:"hummanIds" bson:"hummanIds"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Human struct {
	ID        primitive.ObjectID `json:"_id" bson:"id"`
	Name      string
	Sex       string
	Type      string
	IsLeader  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GroupsDeleteRequest struct {
	ID string `json:"_id"`
}
type GroupDeleteRequest struct {
	ID string `json:"_id"`
}
type HumanDeleteRequest struct {
	ID string `json:"_id"`
}
