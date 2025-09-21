package inquire_password_update

import (
	"vdm/core/dependencies/mailer"
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/inquire"
	Method = fiber.MethodPost
)

func Route(cfg env.SecurityConfig, clientURL string, db *gorm.DB, mailer mailer.Mailer) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:                      repo,
		mailer:                    mailer,
		clientURL:                 clientURL,
		passwordUpdateTokenSecret: cfg.PasswordUpdateTokenSecret,
		passwordUpdateTokenTTL:    cfg.PasswordUpdateTokenTTL,
	}
	handler := &handler{svc}

	return fiberx.NewRoute(Method, Path, handler.inquirePasswordUpdate)
}
