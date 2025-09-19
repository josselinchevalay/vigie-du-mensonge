package set_auth_cookies

import (
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

func Middleware() *fiberx.Middleware {
	return fiberx.NewMiddleware(func(c *fiber.Ctx) error {
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

		isProd := env.Config.ActiveProfile == "prod"

		var sameSite string
		if isProd {
			sameSite = "Strict"
		} else {
			sameSite = "Lax"
		}

		c.Cookie(&fiber.Cookie{
			Name:     local_keys.AccessToken,
			Value:    accessToken.Token,
			Expires:  accessToken.Expiry,
			SameSite: sameSite,
			Secure:   isProd,
			HTTPOnly: true,
			Path:     "/",
		})

		c.Cookie(&fiber.Cookie{
			Name:     local_keys.RefreshToken,
			Value:    refreshToken.Token.String(),
			Expires:  refreshToken.Expiry,
			SameSite: sameSite,
			Secure:   isProd,
			HTTPOnly: true,
			Path:     "/",
		})

		return nil
	})
}
