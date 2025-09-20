package set_auth_cookies

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	setAuthCookies(c *fiber.Ctx) error
}

type handler struct {
	isProd   bool
	sameSite string
}

func (h *handler) setAuthCookies(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		return err
	}

	accessToken, ok := c.Locals(local_keys.AccessToken).(locals.AccessToken)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "access token not found"}
	}

	refreshToken, ok := c.Locals(local_keys.RefreshToken).(locals.RefreshToken)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "refresh token not found"}
	}

	c.Cookie(&fiber.Cookie{
		Name:     local_keys.AccessToken,
		Value:    accessToken.Token,
		Expires:  accessToken.Expiry,
		SameSite: h.sameSite,
		Secure:   h.isProd,
		HTTPOnly: true,
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     local_keys.RefreshToken,
		Value:    refreshToken.Token.String(),
		Expires:  refreshToken.Expiry,
		SameSite: h.sameSite,
		Secure:   h.isProd,
		HTTPOnly: true,
		Path:     "/",
	})

	return nil
}
