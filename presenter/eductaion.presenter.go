package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Groups struct {
	ID        primitive.ObjectID   `json:"_id" bson:"id"`
	Title     string               `json:"title" bson:"title"`
	IsPublic  bool                 `json:"isPublic" bson:"isPublic"`
	Groups    []primitive.ObjectID `json:"groups" bson:"groups"`
	Creator   primitive.ObjectID   `json:"creator" bson:"creator"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}
type Group struct {
	ID        primitive.ObjectID   `json:"_id" bson:"id"`
	Name      string               `json:"name" bson:"name"`
	Place     string               `json:"place" bson:"place"`
	Type      string               `json:"type" bson:"type"`
	HumanIds  []primitive.ObjectID `json:"hummanIds" bson:"hummanIds"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt" bson:"updatedAt"`
}
type Human struct {
	ID        primitive.ObjectID `json:"_id" bson:"id"`
	Name      string             `json:"name" bson:"name"`
	Sex       string             `json:"sex" bson:"sex"`
	Type      string             `json:"type" bson:"type"`
	IsLeader  bool               `json:"isLeader" bson:"isLeader"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func GroupsSuccessResponse(data *entities.Groups) *fiber.Map {
	groups := Groups{
		ID:        data.ID,
		Title:     data.Title,
		IsPublic:  data.IsPublic,
		Groups:    data.Groups,
		Creator:   data.Creator,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"data":   groups,
		"error":  nil,
	}
}

func GroupsListSuccessResponse(data *[]Groups) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func GroupSuccessRepsonse(data *entities.Group) *fiber.Map {
	group := Group{
		ID:        data.ID,
		Name:      data.Name,
		Place:     data.Place,
		Type:      data.Type,
		HumanIds:  data.HumanIds,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return &fiber.Map{
		"status": true,
		"data":   group,
		"error":  nil,
	}
}

func GroupListSuccessResponse(data *[]Group) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func HumanSuccessResponse(data *entities.Human) *fiber.Map {
	human := Human{
		ID:        data.ID,
		Name:      data.Name,
		Sex:       data.Sex,
		Type:      data.Type,
		IsLeader:  data.IsLeader,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   human,
		"error":  nil,
	}

}
func PeopleSuccessResponse(data *[]Human) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func EducationErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
