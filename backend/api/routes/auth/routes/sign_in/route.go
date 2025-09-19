package sign_in

import (
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/sign-in"
	Method = fiber.MethodPost
)

func Route(db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:              repo,
		accessTokenSecret: env.Config.Security.AccessTokenSecret,
		accessTokenTTL:    env.Config.Security.AccessTokenTTL,
		refreshTokenTTL:   env.Config.Security.RefreshTokenTTL,
	}
	handler := &handler{svc}

	return fiberx.NewRoute(Method, Path, handler.signIn)
}
