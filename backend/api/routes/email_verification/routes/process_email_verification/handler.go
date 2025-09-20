package process_email_verification

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	processEmailVerification(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

type processRequest struct {
	Token string `json:"token" validate:"required"`
}

func (h *handler) processEmailVerification(c *fiber.Ctx) error {
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to retrieve authenticated user from locals"}
	}

	var req processRequest
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(req); err != nil {
		return err
	}

	if err := h.svc.processEmailVerification(authedUser, req.Token); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
