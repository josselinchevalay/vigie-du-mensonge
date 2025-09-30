package moderator_find_article

import (
	"fmt"
	"vdm/core/dto/response_dto"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	findArticlesByReferenceForModerator(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) findArticlesByReferenceForModerator(c *fiber.Ctx) error {
	articleRef, err := uuid.Parse(c.Params(local_keys.ArticleReference))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article reference"}
	}

	articles, err := h.repo.findArticlesByReference(articleRef)
	if err != nil {
		return fmt.Errorf("failed to find articles by reference for moderator: %v", err)
	}

	if len(articles) == 0 {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: fmt.Sprintf("article with reference %s not found", articleRef)}
	}

	resDTO := make([]response_dto.Article, len(articles))
	for i := range articles {
		resDTO[i] = response_dto.NewArticle(articles[i])
	}

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
