package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weekly struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}

func WeeklySucccessResponse(data *entities.Weekly) *fiber.Map {
	weekly := Weekly{
		ID:    data.ID,
		Title: data.Title,
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
