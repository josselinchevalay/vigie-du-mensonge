package sign_in

import (
	"vdm/core/locals/local_keys"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	signIn(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) signIn(c *fiber.Ctx) error {
	var req RequestDTO
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(req); err != nil {
		return err
	}

	user, accessToken, refreshToken, err := h.svc.signIn(req)
	if err != nil {
		return err
	}

	c.Locals(local_keys.AccessToken, accessToken)
	c.Locals(local_keys.RefreshToken, refreshToken)

	return c.Status(fiber.StatusOK).JSON(ResponseDTO{
		AccessTokenExpiry:  accessToken.Expiry,
		RefreshTokenExpiry: refreshToken.Expiry,
		EmailVerified:      user.EmailVerified,
		Roles:              user.RoleNames(),
	})
}
