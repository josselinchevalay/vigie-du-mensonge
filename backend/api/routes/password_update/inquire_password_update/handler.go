package inquire_password_update

import (
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	inquirePasswordUpdate(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) inquirePasswordUpdate(c *fiber.Ctx) error {
	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(reqDTO); err != nil {
		return err
	}

	if err := h.svc.inquirePasswordUpdate(reqDTO.Email); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
