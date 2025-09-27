package get_pending_articles

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/pending"
	Method = fiber.MethodGet
)

/*

We consider an article as "pending" if article.Status == UNDER_REVIEW && article.ModeratorID == nil

We'll consider adding a PENDING status to the article model in the future.

*/

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	handler := &handler{repo}
	return fiberx.NewRoute(Method, Path, handler.getPendingArticlesForModerator)
}
