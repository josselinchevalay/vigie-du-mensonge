package auth

import (
	"vdm/api/routes/auth/middlewares/set_auth_cookies"
	"vdm/api/routes/auth/routes/refresh"
	"vdm/api/routes/auth/routes/sign_in"
	"vdm/api/routes/auth/routes/sign_out"
	"vdm/api/routes/auth/routes/sign_up"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/auth"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		set_auth_cookies.Middleware(),

		sign_up.Route(deps.GormDB()),
		sign_in.Route(deps.GormDB()),
		refresh.Route(deps.GormDB()),
		sign_out.Route(deps.GormDB()),
	)

	return group
}
