package refresh

import (
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/refresh"
	Method = fiber.MethodPost
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:              repo,
		accessTokenTTL:    env.Config.Security.AccessTokenTTL,
		refreshTokenTTL:   env.Config.Security.RefreshTokenTTL,
		accessTokenSecret: env.Config.Security.AccessTokenSecret,
	}
	handler := &handler{svc}

	return fiberx.NewRoute(Method, Path, handler.refresh)
}
