package validate_csrf

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

func Middleware(csrfTokenSecret []byte) *fiberx.Middleware {
	handler := &handler{csrfTokenSecret}

	return fiberx.NewMiddleware(
		skip.New(
			handler.validateCsrf,
			func(c *fiber.Ctx) bool {
				method := c.Method()

				return method != fiber.MethodPost &&
					method != fiber.MethodPut &&
					method != fiber.MethodPatch &&
					method != fiber.MethodDelete
			},
		))
}
