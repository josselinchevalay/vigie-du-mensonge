package find_redactor_article_details

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/details"
	Method = fiber.MethodGet
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	handler := &handler{repo}
	return fiberx.NewRoute(Method, Path, handler.findArticleDetailsForRedactor)
}
