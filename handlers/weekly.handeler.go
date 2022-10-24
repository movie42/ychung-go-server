package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/movie42/ychung-go-server/pkg/service"
	"github.com/movie42/ychung-go-server/presenter"
)

func GetWeekies(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchWeeklies()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}
		return c.JSON(presenter.WeekliesSuccessResponse(fetched))
	}
}

func CreateWeekly(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}

func UpdateWeekly(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}

func DeleteWeekly(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
