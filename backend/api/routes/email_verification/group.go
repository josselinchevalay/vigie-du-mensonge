package email_verification

import (
	"vdm/api/routes/email_verification/routes/inquire_email_verification"
	"vdm/api/routes/email_verification/routes/process_email_verification"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/email-verification"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		inquire_email_verification.Route(deps.Config.Security, deps.Config.ClientURL, deps.Mailer),
		process_email_verification.Route(deps.Config.Security, deps.GormDB()),
	)

	return group
}
