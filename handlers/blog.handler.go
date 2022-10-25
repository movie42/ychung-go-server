package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/movie42/ychung-go-server/entities"
	"github.com/movie42/ychung-go-server/pkg/service"
	"github.com/movie42/ychung-go-server/presenter"
)

func GetBlogPosts(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchBlogPosts()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		return c.JSON(presenter.BlogPostsSuccessResponse(fetched))
	}
}

func CreateBlogPost(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.BlogPost
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		if requestBody.Title != "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BlogErrorResponse(errors.New("블로그 포스트에는 제목이 반드시 필요합니다.")))
		}
		result, err := service.InsertBlogPost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		return c.JSON(presenter.BlogPostSuccessResponse(result))
	}
}

func UpdateBlogPosts(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.BlogPost
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		result, err := service.UpdateBlogPost(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		return c.JSON(presenter.BlogPostSuccessResponse(result))
	}
}

func DeleteBlogPost(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.BlogPostDeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		postID := requestBody.ID
		err = service.DeleteBlogPost(postID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BlogErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "블로그 포스트가 성공적으로 삭제되었습니다.",
			"error":  nil,
		})
	}
}
