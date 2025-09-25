package update_redactor_article

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/"
	Method = fiber.MethodPut
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	svc := &service{repo}
	h := &handler{svc}
	return fiberx.NewRoute(Method, Path, h.updateArticleForRedactor)
}
