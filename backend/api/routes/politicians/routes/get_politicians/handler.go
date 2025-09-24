package get_politicians

import "github.com/gofiber/fiber/v2"

type Handler interface {
	getPoliticians(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) getPoliticians(c *fiber.Ctx) error {
	respDTO, err := h.svc.getAndMapPoliticians()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(respDTO)
}
