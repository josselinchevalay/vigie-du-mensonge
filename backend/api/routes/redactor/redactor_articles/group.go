package redactor_articles

import (
	"vdm/api/routes/redactor/redactor_articles/redactor_find_article"
	"vdm/api/routes/redactor/redactor_articles/redactor_get_articles"
	"vdm/api/routes/redactor/redactor_articles/redactor_save_article"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		redactor_get_articles.Route(deps.GormDB()),
		redactor_save_article.Route(deps.GormDB()),
		redactor_find_article.Route(deps.GormDB()),
	)

	return group
}
