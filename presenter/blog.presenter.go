package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
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

func BlogPostSuccessResponse(data *entities.BlogPost) *fiber.Map {
	posts := BlogPost{
		ID:        data.ID,
		Title:     data.Title,
		Paragraph: data.Paragraph,
		Creator:   data.Creator,
		Comments:  data.Comments,
		Views:     data.Views,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"data":   posts,
		"error":  nil,
	}
}

func BlogPostsSuccessResponse(data *[]BlogPost) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func BlogErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
