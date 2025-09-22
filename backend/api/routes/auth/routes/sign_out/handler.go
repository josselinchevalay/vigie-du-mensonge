package sign_out

import (
	"time"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	signOut(c *fiber.Ctx) error
}

type handler struct {
	svc              Service
	accessCookieName string
}

func (h *handler) signOut(c *fiber.Ctx) error {
	c.Locals(local_keys.AccessToken, locals.AccessToken{Expiry: time.Now()})
	c.Locals(local_keys.RefreshToken, locals.RefreshToken{Expiry: time.Now()})

	h.svc.signOut(c.Cookies(h.accessCookieName))

	return c.SendStatus(fiber.StatusNoContent)
}
