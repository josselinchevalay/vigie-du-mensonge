package redactor

import (
	"fmt"
	"vdm/api/routes/redactor/redactor_articles"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
)

const Prefix = "/redactor"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		fiberx.NewMiddleware(func(c *fiber.Ctx) error {
			authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
			if !ok {
				return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
			}

			if !authedUser.HasRole(models.RoleRedactor) {
				return &fiber.Error{Code: fiber.StatusForbidden,
					Message: fmt.Sprintf("user %s does not have role %s", authedUser.ID, models.RoleRedactor)}
			}

			return c.Next()
		}),

		redactor_articles.Group(deps),
	)

	return group
}
