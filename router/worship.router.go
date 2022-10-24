package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/handlers"
	"github.com/movie42/ychung-go-server/pkg/weekly"
)

func WorshipRouter(app fiber.Router, service weekly.Service) {
	app.Get("/", handlers.GetWeekies(service))
}
