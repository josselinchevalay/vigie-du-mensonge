package email_verification

import (
	"vdm/api/routes/email_verification/routes/initiate_email_verification"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/email-verification"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		initiate_email_verification.Route(deps.Config.Security, deps.Mailer),
	)

	return group
}
