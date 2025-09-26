package save_redactor_draft

import (
	"vdm/core/locals"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	saveDraftArticleForRedactor(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) saveDraftArticleForRedactor(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(reqDTO); err != nil {
		return err
	}

	article, err := reqDTO.toArticle(authedUser.ID)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	if err = h.repo.saveArticle(&article); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
