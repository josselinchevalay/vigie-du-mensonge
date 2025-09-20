package email_verification

import (
	"vdm/core/fiberx"
)

const Prefix = "/email-verification"

func Group() *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add()

	return group
}
