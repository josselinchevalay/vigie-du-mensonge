package get_politicians

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
	svc := &service{repo}
	handler := &handler{svc}

	group := fiberx.NewGroup(Path)

	group.Add(
		fiberx.NewMiddleware(cache.New(cache.Config{
			Expiration: 24 * time.Hour,
		})),

		fiberx.NewRoute(Method, Path, handler.getPoliticians),
	)

	return group
}
