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

func Route(db *gorm.DB, cfg env.SecurityConfig) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:              repo,
		accessTokenTTL:    cfg.AccessTokenTTL,
		refreshTokenTTL:   cfg.RefreshTokenTTL,
		accessTokenSecret: cfg.AccessTokenSecret,
	}
	handler := &handler{
		svc:               svc,
		refreshCookieName: cfg.RefreshCookieName,
	}

	return fiberx.NewRoute(Method, Path, handler.refresh)
}
