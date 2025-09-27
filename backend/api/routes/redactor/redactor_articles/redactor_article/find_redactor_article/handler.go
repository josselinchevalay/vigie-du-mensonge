package find_redactor_article

import (
	"vdm/core/dto/response_dto"
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

	articleRef, err := uuid.Parse(c.Params(local_keys.ArticleReference))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article reference"}
	}

	article, otherVersions, err := h.repo.findRedactorArticlesByReference(authedUser.ID, articleRef)
	if err != nil {
		return err
	}

	resDTO := response_dto.NewArticle(article, otherVersions)

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
