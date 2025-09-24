package process_sign_up

import (
	"vdm/core/locals/local_keys"
	"vdm/core/models"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	processSignUp(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) processSignUp(c *fiber.Ctx) error {
	var req RequestDTO
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(req); err != nil {
		return err
	}

	accessToken, refreshToken, err := h.svc.createUserAndBuildTokens(req)
	if err != nil {
		return err
	}

	c.Locals(local_keys.AccessToken, accessToken)
	c.Locals(local_keys.RefreshToken, refreshToken)

	return c.Status(fiber.StatusCreated).JSON(ResponseDTO{
		AccessTokenExpiry:  accessToken.Expiry,
		RefreshTokenExpiry: refreshToken.Expiry,
		Roles:              []models.RoleName{models.RoleRedactor},
	})
}
