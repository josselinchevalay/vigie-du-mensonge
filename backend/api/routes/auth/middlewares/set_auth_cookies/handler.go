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
	cookieSecure      bool
	cookieSameSite    string
	accessCookieName  string
	refreshCookieName string
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
		Name:     h.accessCookieName,
		Value:    accessToken.Token,
		Expires:  accessToken.Expiry,
		SameSite: h.cookieSameSite,
		Secure:   h.cookieSecure,
		HTTPOnly: true,
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     h.refreshCookieName,
		Value:    refreshToken.Token.String(),
		Expires:  refreshToken.Expiry,
		SameSite: h.cookieSameSite,
		Secure:   h.cookieSecure,
		HTTPOnly: true,
		Path:     "/",
	})

	return nil
}
