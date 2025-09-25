package create_draft_article

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/"
	Method = fiber.MethodPost
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repostiroy{db}
	svc := &service{repo}
	handler := &handler{svc}

	return fiberx.NewRoute(Method, Path, handler.createDraftArticle)
}
