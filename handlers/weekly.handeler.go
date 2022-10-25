package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/movie42/ychung-go-server/entities"
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
		var requestBody entities.Weekly
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}
		if requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(errors.New("주보 제목을 반드시 써야합니다.")))
		}

		result, err := service.InsertWeekly(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}
		return c.JSON(presenter.WeeklySucccessResponse(result))
	}
}

func UpdateWeekly(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Weekly
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}

		result, err := service.UpdateWeekly(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}

		return c.JSON(presenter.WeeklySucccessResponse(result))

	}
}

func DeleteWeekly(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.WeeklyDeleteRequest

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}

		weeklyID := requestBody.ID
		err = service.DeleteWeekly(weeklyID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.WeekliesErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "게시물이 성공적으로 삭제되었습니다.",
			"err":    nil,
		})
	}
}
