package admin

import (
	"fmt"
	"vdm/api/routes/admin/admin_users"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
)

const Prefix = "/admin"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		fiberx.NewMiddleware(func(c *fiber.Ctx) error {
			authedUser, ok := c.Locals(local_keys.AuthedUser).(locals.AuthedUser)
			if !ok {
				return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
			}

			if !authedUser.HasRole(models.RoleAdmin) {
				return &fiber.Error{Code: fiber.StatusForbidden,
					Message: fmt.Sprintf("user %s does not have role %s", authedUser.ID, models.RoleModerator)}
			}

			return c.Next()
		}),

		admin_users.Group(deps),
	)

	return group
}
