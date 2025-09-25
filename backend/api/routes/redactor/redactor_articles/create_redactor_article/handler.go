package create_redactor_article

import (
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	createRedactorArticle(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) createRedactorArticle(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := reqDTO.Validate(); err != nil {
		return err
	}

	articleID, err := h.svc.mapAndCreateArticle(authedUser.ID, reqDTO)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"articleID": articleID,
	})
}
