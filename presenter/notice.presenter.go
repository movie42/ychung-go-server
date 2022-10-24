package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Title     string               `json:"title" bson:"title"`
	IsWeekly  bool                 `json:"isWeekly" bson:"isWeekly,omitempty"`
	StartDate string               `json:"startDate" bson:"startDate,omitempty"`
	EndDate   string               `json:"endDate" bson:"endDate,omitempty"`
	Summary   string               `json:"summary" bson:"summary,omitempty"`
	Paragraph string               `json:"paragraph" bson:"paragraph"`
	Creator   primitive.ObjectID   `json:"creator" bson:"creator"`
	Comments  []primitive.ObjectID `json:"comments" bson:"comments"`
	Views     int32                `json:"views" bson:"views"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}

func NoticeSuccessResponse(data *entities.Notice) *fiber.Map {
	notice := Notice{
		ID:        data.ID,
		Title:     data.Title,
		IsWeekly:  data.IsWeekly,
		Paragraph: data.Paragraph,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Summary:   data.Summary,
		Creator:   data.Creator,
		Comments:  data.Comments,
		Views:     data.Views,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"data":   notice,
		"error":  nil,
	}
}

func NoticesSuccessResponse(data *[]Notice) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func NoticeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
