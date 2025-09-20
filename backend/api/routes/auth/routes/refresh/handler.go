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
	svc Service
}

func (h *handler) refresh(c *fiber.Ctx) error {
	rftID, err := uuid.Parse(c.Cookies(local_keys.RefreshToken))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid refresh token"}
	}

	user, accessToken, refreshToken, err := h.svc.refresh(rftID)
	if err != nil {
		return err
	}

	c.Locals(local_keys.AccessToken, accessToken)
	c.Locals(local_keys.RefreshToken, refreshToken)

	return c.Status(fiber.StatusOK).JSON(ResponseDTO{
		AccessTokenExpiry:  accessToken.Expiry,
		RefreshTokenExpiry: refreshToken.Expiry,
		EmailVerified:      user.EmailVerified,
		Roles:              user.MapRoles(),
	})
}
