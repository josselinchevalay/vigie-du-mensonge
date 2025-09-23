package inquire_sign_up

import (
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	inquireSignUp(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) inquireSignUp(c *fiber.Ctx) error {
	var req RequestDTO
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(req); err != nil {
		return err
	}

	if err := h.svc.checkUserAndSendEmail(req.Email); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
