package router

import (
	"github.com/gofiber/fiber/v2"
)

func NoticeRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("it is noitce Router!!!")
	})
}
