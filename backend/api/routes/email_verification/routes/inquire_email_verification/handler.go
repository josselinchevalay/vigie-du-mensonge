package inquire_email_verification

import (
	"errors"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	inquireEmailVerification(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) inquireEmailVerification(c *fiber.Ctx) error {
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return errors.New("failed to retrieve authenticated user from locals")
	}

	if authedUser.EmailVerified {
		return &fiber.Error{Code: fiber.StatusConflict, Message: "email already verified"}
	}

	if err := h.svc.sendEmailAndCreateToken(authedUser); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
