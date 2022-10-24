package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/pkg/notice"
	"github.com/movie42/ychung-go-server/presenter"
)

func UpdateNotice(service notice.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Notice
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		result, err := service.UpdateNotice(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		return c.JSON(presenter.NoticeSuccessResponse(result))
	}
}

func DeleteNotice(service notice.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.NoticeDeleteRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		noticeID := requestBody.ID
		err = service.DeleteNotice(noticeID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "게시물이 성공적으로 삭제되었습니다.",
			"err":    nil,
		})
	}
}

func AddNotice(service notice.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Notice
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		if requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.NoticeErrorResponse(errors.New("제목이 꼭 필요합니다.")))
		}
		result, err := service.InsertNotice(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		return c.JSON(presenter.NoticeSuccessResponse(result))
	}
}

func GetNotices(service notice.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchNotices()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.NoticeErrorResponse(err))
		}
		return c.JSON(presenter.NoticesSuccessResponse(fetched))
	}
}
