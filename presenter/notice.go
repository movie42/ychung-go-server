package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notice struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title"`
}

func NoticeSuccessResponse(data *entities.Notice) *fiber.Map {
	notice := Notice{
		ID:    data.ID,
		Title: data.Title,
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
