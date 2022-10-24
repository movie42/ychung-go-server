package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/pkg/weekly"
	"github.com/movie42/ychung-go-server/presenter"
)

func GetWeekies(service weekly.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchWeeklies()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}
		return c.JSON(presenter.WeekliesSuccessResponse(fetched))
	}
}
