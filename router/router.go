package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/handlers"
	"github.com/movie42/ychung-go-server/pkg/service"
)

func NoticeRouter(app fiber.Router, service service.Service) {
	notice := app.Group("/notice")

	notice.Get("/", handlers.GetNotices(service))
	notice.Post("/", handlers.AddNotice(service))
	notice.Delete("/", handlers.DeleteNotice(service))
	notice.Put("/", handlers.UpdateNotice(service))
}

func WorshipRouter(app fiber.Router, service service.Service) {
	worship := app.Group("/worship")

	worship.Get("/weekly", handlers.GetWeekies(service))
	worship.Post("/weekly", handlers.CreateWeekly(service))
	worship.Put("/weekly", handlers.UpdateWeekly(service))
	worship.Delete("/weekly", handlers.DeleteWeekly(service))
}

func BlogRouter(app fiber.Router, service service.Service) {
	blog := app.Group("/blog")

	blog.Get("/", handlers.GetBlogPosts(service))
	blog.Post("/", handlers.CreateBlogPost(service))
	blog.Put("/", handlers.UpdateBlogPosts(service))
	blog.Delete("/", handlers.DeleteBlogPost(service))
}

func EducationRouter(app fiber.Router, service service.Service) {
	// education := app.Group("/education")

	// education.Get("/")
}
