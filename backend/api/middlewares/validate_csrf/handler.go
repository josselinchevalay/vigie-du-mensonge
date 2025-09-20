package validate_csrf

import (
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	validateCsrf(c *fiber.Ctx) error
}

type handler struct {
	csrfTokenSecret []byte
}

func (h *handler) validateCsrf(c *fiber.Ctx) error {
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to retrieve authenticated user from locals"}
	}

	token := c.Get("csrf-token")

	if tokenUser, err := jwt_utils.ParseJWT(token, h.csrfTokenSecret); err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	} else if tokenUser.ID != authedUser.ID {
		return &fiber.Error{Code: fiber.StatusForbidden, Message: "credentials do not match"}
	}

	return c.Next()
}
