package inquire_password_update

import (
	"vdm/core/dependencies/mailer"
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/inquire"
	Method = fiber.MethodPost
)

func Route(cfg env.SecurityConfig, clientURL string, mailer mailer.Mailer) *fiberx.Route {
	handler := &handler{
		passwordUpdateTokenSecret: cfg.PasswordUpdateTokenSecret,
		passwordUpdateTokenTTL:    cfg.PasswordUpdateTokenTTL,
		clientURL:                 clientURL,
		mailer:                    mailer,
	}

	return fiberx.NewRoute(Method, Path, handler.inquirePasswordUpdate)
}
