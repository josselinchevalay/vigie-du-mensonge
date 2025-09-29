package get_published_articles

import (
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
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
		// TODO: enable caching in prod
		//fiberx.NewMiddleware(cache.New(cache.Config{
		//	CacheControl: true,
		//	Expiration:   24 * time.Hour,
		//})),

		fiberx.NewRoute(Method, Path, handler.getPublishedArticles),
	)

	return group
}
