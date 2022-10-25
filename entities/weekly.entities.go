package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weekly struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Title         string             `json:"title" bson:"title"`
	Word          string             `json:"word" bson:"word"`
	Chapter       int8               `json:"chapter" bson:"chapter"`
	Verse         int8               `json:"verse" bson:"verse"`
	Verse_End     int8               `json:"verse_end" bson:"verse_end"`
	Pastor        string             `json:"pastor" bson:"pastor"`
	Prayer        string             `json:"prayer" bson:"prayer"`
	Advertisement string             `json:"advertisement" bson:"advertisement"`
	Reader        string             `json:"reader" bson:"reader"`
	Offering      string             `json:"offering" bson:"offering"`
	Benediction   string             `json:"benediction" bson:"benediction"`
	Creator       primitive.ObjectID `json:"creator" bson:"creator"`
	Views         int32              `json:"views" bson:"views"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type WeeklyDeleteRequest struct {
	ID string `json:"id"`
}
