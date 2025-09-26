package find_redactor_article

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	findArticleDetailsForRedactor(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) findArticleDetailsForRedactor(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articleID, err := uuid.Parse(c.Params(local_keys.ArticleID))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article id"}
	}

	article, err := h.repo.findRedactorArticle(articleID, authedUser.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(newArticleDTO(article))
}
