package inquire_email_verification

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
	h := &handler{
		emailVerificationTokenSecret: cfg.EmailVerificationTokenSecret,
		emailVerificationTokenTTL:    cfg.EmailVerificationTokenTTL,
		mailer:                       mailer,
		clientURL:                    clientURL,
	}
	return fiberx.NewRoute(Method, Path, h.inquireEmailVerification)
}
