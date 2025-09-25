package redactor_articles

import (
	"vdm/api/routes/redactor/redactor_articles/create_draft_article"
	"vdm/api/routes/redactor/redactor_articles/get_redactor_articles"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_redactor_articles.Route(deps.GormDB()),
		create_draft_article.Route(deps.GormDB()),
	)

	return group
}
