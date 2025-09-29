package find_published_article

import (
	"fmt"
	"vdm/core/dto/response_dto"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	findPublishedArticleForUser(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) findPublishedArticleForUser(c *fiber.Ctx) error {
	articleID, err := uuid.Parse(c.Params(local_keys.ArticleID))

	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article id"}
	}

	article, err := h.repo.findPublishedArticle(articleID)
	if err != nil {
		return err
	}
	if article == nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: fmt.Sprintf("article with id %s not found", articleID)}
	}

	return c.Status(fiber.StatusOK).JSON(response_dto.NewArticle(*article))
}
