package get_csrf

import (
	"time"
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

const (
	Path   = "/csrf-token"
	Method = fiber.MethodGet
)

func Group() *fiberx.Group {
	//setup specific group for custom rate limiter on /csrf-token
	group := fiberx.NewGroup(Path)

	group.Add(
		fiberx.NewMiddleware(limiter.New(limiter.Config{
			Max:        50,
			Expiration: 30 * time.Second,
		})),

		fiberx.NewRoute(Method, "/", func(c *fiber.Ctx) error {
			token := c.Locals(local_keys.CsrfToken)

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"csrfToken": token,
			})
		}),
	)

	return group
}
