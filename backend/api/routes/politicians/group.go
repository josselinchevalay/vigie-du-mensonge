package politicians

import (
	"vdm/api/routes/politicians/routes/get_politicians"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/politicians"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_politicians.Group(deps.GormDB()),
	)

	return group
}
