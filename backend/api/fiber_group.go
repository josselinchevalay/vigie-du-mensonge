package api

import (
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/api/v1"

func FiberGroup(deps *dependencies.Dependencies) *fiberx.Group {
	fiberGroup := fiberx.NewGroup(Prefix)

	fiberGroup.Add(
	//TODO: impl api
	)

	return fiberGroup
}
