package articles

import (
	"vdm/api/routes/articles/find_published_article"
	"vdm/api/routes/articles/get_published_articles"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_published_articles.Group(deps.GormDB()),
		find_published_article.Route(deps.GormDB()),
	)

	return group
}
