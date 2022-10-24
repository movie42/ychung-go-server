package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/handlers"
	"github.com/movie42/ychung-go-server/pkg/service"
)

func WorshipRouter(app fiber.Router, service service.Service) {
	app.Get("/weeklies", handlers.GetWeekies(service))
}
