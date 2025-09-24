package create_draft_article

import (
	"vdm/core/locals"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	createDraftArticle(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) createDraftArticle(c *fiber.Ctx) error {
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

	if err := h.svc.mapAndCreateArticle(authedUser.ID, reqDTO); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
