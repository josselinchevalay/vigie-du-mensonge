package save_moderator_review

import (
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/:" + local_keys.ArticleID + "/review"
	Method = fiber.MethodPost
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	handler := &handler{repo}
	return fiberx.NewRoute(Method, Path, handler.saveArticleReviewForModerator)
}
