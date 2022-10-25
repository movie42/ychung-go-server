package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weekly struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
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

func WeeklySucccessResponse(data *entities.Weekly) *fiber.Map {
	weekly := Weekly{
		ID:            data.ID,
		Title:         data.Title,
		Word:          data.Word,
		Chapter:       int8(data.Chapter),
		Verse:         int8(data.Verse),
		Verse_End:     int8(data.Verse_End),
		Pastor:        data.Pastor,
		Prayer:        data.Prayer,
		Advertisement: data.Advertisement,
		Reader:        data.Reader,
		Offering:      data.Offering,
		Benediction:   data.Benediction,
		Creator:       data.Creator,
		Views:         data.Views,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"data":   weekly,
		"error":  nil,
	}
}

func WeekliesSuccessResponse(data *[]Weekly) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func WeekliesErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
