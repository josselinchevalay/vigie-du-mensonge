package get_redactor_articles

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/"
	Method = fiber.MethodGet
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	svc := &service{repo}
	handler := &handler{svc}
	return fiberx.NewRoute(Method, Path, handler.getRedactorArticles)
}
