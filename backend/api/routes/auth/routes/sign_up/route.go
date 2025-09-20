package sign_up

import (
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/sign-up"
	Method = fiber.MethodPost
)

func Route(db *gorm.DB, cfg env.SecurityConfig) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:              repo,
		accessTokenSecret: cfg.AccessTokenSecret,
		accessTokenTTL:    cfg.AccessTokenTTL,
		refreshTokenTTL:   cfg.RefreshTokenTTL,
	}
	handler := &handler{svc: svc}

	return fiberx.NewRoute(Method, Path, handler.signUp)
}
