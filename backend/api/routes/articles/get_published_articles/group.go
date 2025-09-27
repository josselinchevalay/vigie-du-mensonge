package get_published_articles

import (
	"time"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"gorm.io/gorm"
)

const (
	Path   = "/"
	Method = fiber.MethodGet
)

func Group(db *gorm.DB) *fiberx.Group {
	repo := &repository{db}
	handler := &handler{repo}

	group := fiberx.NewGroup(Path)

	group.Add(
		fiberx.NewMiddleware(cache.New(cache.Config{
			CacheControl: true,
			Expiration:   24 * time.Hour,
		})),

		fiberx.NewRoute(Method, Path, handler.getPublishedArticles),
	)

	return group
}
