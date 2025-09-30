package moderator_articles

import (
	"vdm/api/routes/moderator/moderator_articles/moderator_claim_article"
	"vdm/api/routes/moderator/moderator_articles/moderator_find_article"
	"vdm/api/routes/moderator/moderator_articles/moderator_get_claimed_articles"
	"vdm/api/routes/moderator/moderator_articles/moderator_get_pending_articles"
	"vdm/api/routes/moderator/moderator_articles/moderator_save_review"
	"vdm/core/dependencies"
	"vdm/core/fiberx"
)

const Prefix = "/articles"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		moderator_get_claimed_articles.Route(deps.GormDB()),
		moderator_get_pending_articles.Route(deps.GormDB()),
		moderator_find_article.Route(deps.GormDB()),
		moderator_claim_article.Route(deps.GormDB()),
		moderator_save_review.Route(deps.GormDB()),
	)

	return group
}
