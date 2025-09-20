package process_password_update

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	processPasswordUpdate(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) processPasswordUpdate(c *fiber.Ctx) error {
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to retrieve authenticated user from locals"}
	}

	var req RequestDTO
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(req); err != nil {
		return err
	}

	if err := h.svc.processPasswordUpdate(authedUser, req.Token, req.NewPassword); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
