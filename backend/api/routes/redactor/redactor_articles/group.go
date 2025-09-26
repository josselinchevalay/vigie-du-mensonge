package redactor_articles

import (
	"vdm/api/routes/redactor/redactor_articles/get_redactor_articles"
	"vdm/api/routes/redactor/redactor_articles/redactor_article"
	"vdm/api/routes/redactor/redactor_articles/save_redactor_article"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_redactor_articles.Route(deps.GormDB()),
		save_redactor_article.Route(deps.GormDB()),

		redactor_article.Group(deps),
	)

	return group
}
