package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/handlers"
	"github.com/movie42/ychung-go-server/pkg/notice"
)

func ApiRouter(app fiber.Router, service notice.Service) {
	app.Get("/notices", handlers.GetNotices(service))
	// api := fiber.New()
	// NoticeRouter(api.Group("/notice"))
	// app.Mount("/api", api)
}
