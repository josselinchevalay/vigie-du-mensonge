package redactor_article

import (
	"vdm/api/routes/redactor/redactor_articles/redactor_article/find_redactor_article"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"
)

const Prefix = "/:" + local_keys.ArticleID

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		find_redactor_article.Route(deps.GormDB()),
	)

	return group
}
