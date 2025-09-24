package me

import (
	"vdm/api/middlewares/locals_authed_user"
	"vdm/api/routes/me/create_draft_article"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/me"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		locals_authed_user.Middleware(deps.Config.Security),
		create_draft_article.Route(deps.GormDB()),
	)

	return group
}
