package api

import (
	"vdm/api/middlewares/cors"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/api/routes/auth"
	"vdm/api/routes/email_verification"
	"vdm/api/routes/password_update"
	"vdm/core/dependencies"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const Prefix = "/api/v1"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		cors.Middleware(deps.Config.AllowOrigins),

		auth.Group(deps),
		password_update.Group(deps),

		locals_authed_user.Middleware(deps.Config.Security),

		email_verification.Group(deps),
	)

	return group
}

func Register(router fiber.Router, deps *dependencies.Dependencies) {
	Group(deps).Register(router)
}
