package password_update

import (
	"vdm/api/routes/password_update/inquire_password_update"
	"vdm/api/routes/password_update/process_password_update"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const (
	Prefix = "/password-update"
)

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		inquire_password_update.Route(deps.Config.Security, deps.Config.ClientURL, deps.Mailer),
		process_password_update.Route(deps.GormDB(), deps.Config.Security),
	)

	return group
}
