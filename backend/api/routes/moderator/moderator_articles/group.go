package moderator_articles

import (
	"vdm/api/routes/moderator/moderator_articles/get_moderator_articles"
	"vdm/api/routes/moderator/moderator_articles/get_pending_articles"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_moderator_articles.Route(deps.GormDB()),
		get_pending_articles.Route(deps.GormDB()),
	)

	return group
}
