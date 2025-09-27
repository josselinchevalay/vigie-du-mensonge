package get_published_articles

import (
	"fmt"
	"vdm/core/dto/response_dto"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	getPublishedArticles(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) getPublishedArticles(c *fiber.Ctx) error {
	articles, err := h.repo.getPublishedArticles()
	if err != nil {
		return fmt.Errorf("failed to get published articles: %v", err)
	}

	respDTO := make([]response_dto.Article, len(articles))

	for i := range articles {
		respDTO[i] = response_dto.NewArticle(articles[i])
	}

	return c.Status(fiber.StatusOK).JSON(respDTO)
}
