package save_redactor_article

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
	repo := &repository{db}
	svc := &service{repo}
	handler := &handler{svc}
	return fiberx.NewRoute(Method, Path, handler.saveArticleForRedactor)
}
