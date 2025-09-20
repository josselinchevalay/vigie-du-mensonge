package get_csrf

import (
	"fmt"
	"time"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	getCsrf(c *fiber.Ctx) error
}

type handler struct {
	csrfTokenSecret []byte
	csrfTokenTTL    time.Duration
}

func (h *handler) getCsrf(c *fiber.Ctx) error {
	authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to retrieve authenticated user from locals"}
	}

	csrfToken, err := jwt_utils.GenerateJWT(authedUser, h.csrfTokenSecret, time.Now().Add(h.csrfTokenTTL))
	if err != nil {
		return fmt.Errorf("failed to generate CSRF token: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"csrfToken": csrfToken,
	})
}
