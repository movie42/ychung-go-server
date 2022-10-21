package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/pkg/notice"
	"github.com/movie42/ychung-go-server/presenter"
)

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
