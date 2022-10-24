package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	ID        primitive.ObjectID   `json:"id"  bson:"_id"`
	Title     string               `json:"title" bson:"title"`
	StartDate string               `json:"startDate" bson:"startDate"`
	EndDate   string               `json:"endDate" bson:"endDate"`
	Summary   string               `json:"summary" bson:"summary"`
	IsWeekly  bool                 `json:"isWeekly" bson:"isWeekly"`
	Paragraph string               `json:"paragraph" bson:"paragraph"`
	Creator   primitive.ObjectID   `json:"creator" bson:"creator"`
	Comments  []primitive.ObjectID `json:"comments" bson:"comments"`
	Views     int32                `json:"views" bson:"views"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}

type NoticeDeleteRequest struct {
	ID string `json:"id"`
}
