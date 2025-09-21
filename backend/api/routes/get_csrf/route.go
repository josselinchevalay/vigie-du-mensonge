package get_csrf

import (
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/csrf-token"
	Method = fiber.MethodGet
)

func Route() *fiberx.Route {
	return fiberx.NewRoute(Method, Path, func(c *fiber.Ctx) error {
		token := c.Locals(local_keys.CsrfToken)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"csrfToken": token,
		})
	})
}
