package refresh

import (
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	refresh(c *fiber.Ctx) error
}

type handler struct {
	svc               Service
	refreshCookieName string
}

func (h *handler) refresh(c *fiber.Ctx) error {
	rft, err := uuid.Parse(c.Cookies(h.refreshCookieName))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid refresh token"}
	}

	user, accessToken, refreshToken, err := h.svc.refresh(rft)
	if err != nil {
		return err
	}

	c.Locals(local_keys.AccessToken, accessToken)
	c.Locals(local_keys.RefreshToken, refreshToken)

	return c.Status(fiber.StatusOK).JSON(ResponseDTO{
		AccessTokenExpiry:  accessToken.Expiry,
		RefreshTokenExpiry: refreshToken.Expiry,
		Roles:              user.RoleNames(),
		Tag:                user.Tag,
	})
}
