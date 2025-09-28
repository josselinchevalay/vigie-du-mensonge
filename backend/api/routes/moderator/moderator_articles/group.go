package moderator_articles

import (
	"vdm/api/routes/moderator/moderator_articles/claim_moderator_article"
	"vdm/api/routes/moderator/moderator_articles/find_moderator_article"
	"vdm/api/routes/moderator/moderator_articles/get_moderator_articles"
	"vdm/api/routes/moderator/moderator_articles/get_pending_articles"
	"vdm/api/routes/moderator/moderator_articles/review_moderator_article"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		get_moderator_articles.Route(deps.GormDB()),
		get_pending_articles.Route(deps.GormDB()),
		find_moderator_article.Route(deps.GormDB()),
		claim_moderator_article.Route(deps.GormDB()),
		review_moderator_article.Route(deps.GormDB()),
	)

	return group
}
