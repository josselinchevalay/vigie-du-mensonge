package get_articles

import "github.com/gofiber/fiber/v2"

type Handler interface {
	getArticles(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) getArticles(c *fiber.Ctx) error {
	respDTO, err := h.svc.getAndMapArticles()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(respDTO)
}
