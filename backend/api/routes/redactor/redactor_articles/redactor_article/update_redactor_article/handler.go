package update_redactor_article

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	updateArticleForRedactor(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) updateArticleForRedactor(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articleID, err := uuid.Parse(c.Params(local_keys.ArticleID))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article id"}
	}

	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := reqDTO.Validate(); err != nil {
		return err
	}

	if err := h.svc.updateRedactorArticle(authedUser.ID, articleID, reqDTO); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
