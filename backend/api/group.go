package api

import (
	"vdm/api/middlewares/cors"
	"vdm/api/routes/auth"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/api/v1"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		cors.Middleware(),
		auth.Group(deps),
	)

	return group
}
