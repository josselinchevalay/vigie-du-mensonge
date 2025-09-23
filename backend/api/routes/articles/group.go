package articles

import "vdm/core/fiberx"

const Prefix = "/articles"

func Group() *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
	//TODO: impl
	)

	return group
}
