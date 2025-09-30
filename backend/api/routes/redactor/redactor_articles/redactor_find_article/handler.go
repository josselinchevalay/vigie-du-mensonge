package redactor_find_article

import (
	"fmt"
	"vdm/core/dto/response_dto"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	findArticlesByReferenceForRedactor(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) findArticlesByReferenceForRedactor(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articleRef, err := uuid.Parse(c.Params(local_keys.ArticleReference))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article reference"}
	}

	articles, err := h.repo.findRedactorArticlesByReference(authedUser.ID, articleRef)
	if err != nil {
		return fmt.Errorf("failed to find articles by reference for redactor: %v", err)
	}

	if len(articles) == 0 {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: "article not found"}
	}

	resDTO := make([]response_dto.Article, len(articles))
	for i := range articles {
		resDTO[i] = response_dto.NewArticle(articles[i])
	}

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
