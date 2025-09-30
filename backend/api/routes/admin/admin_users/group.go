package admin_users

import (
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/users"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
	//TODO: impl routes
	)

	return group
}
