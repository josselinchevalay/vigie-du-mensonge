package get_pending_articles

import (
	"fmt"
	"vdm/core/dto/response_dto"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	getPendingArticlesForModerator(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) getPendingArticlesForModerator(c *fiber.Ctx) error {
	articles, err := h.repo.getPendingArticles()
	if err != nil {
		return fmt.Errorf("failed to get pending articles: %v", err)
	}

	resDTO := make([]response_dto.Article, len(articles))
	for i := range articles {
		resDTO[i] = response_dto.NewArticle(articles[i])
	}

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
