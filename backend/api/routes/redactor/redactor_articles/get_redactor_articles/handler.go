package get_redactor_articles

import (
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	getRedactorArticles(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) getRedactorArticles(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	resDTO, err := h.svc.getAndMapUserArticles(authedUser.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
