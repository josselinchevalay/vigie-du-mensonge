package api

import (
	"vdm/api/middlewares/dev_cors"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/api/middlewares/validate_csrf"
	"vdm/api/routes/auth"
	"vdm/api/routes/email_verification"
	"vdm/api/routes/get_csrf"
	"vdm/api/routes/password_update"
	"vdm/core/dependencies"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const Prefix = "/api/v1"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	if deps.Config.ActiveProfile != "prod" {
		group.Add(dev_cors.Middleware(deps.Config.AllowOrigins))
	}

	group.Add(
		validate_csrf.Middleware(deps.Config.ActiveProfile == "prod"),
		get_csrf.Route(),

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
