package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/handlers"
	"github.com/movie42/ychung-go-server/pkg/notice"
)

func NoticeRouter(app fiber.Router, service notice.Service) {
	app.Get("/", handlers.GetNotices(service))
	app.Post("/", handlers.AddNotice(service))
	app.Delete("/", handlers.DeleteNotice(service))
	app.Put("/", handlers.UpdateNotice(service))
}
