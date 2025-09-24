package authorize_authed_user

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	authorizeAuthedUser(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) authorizeAuthedUser(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	if err := h.svc.authorizeAuthedUser(&authedUser); err != nil {
		return err
	}

	c.Locals(local_keys.AuthedUser, authedUser)
	return c.Next()
}
