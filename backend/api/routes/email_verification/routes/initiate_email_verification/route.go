package initiate_email_verification

import (
	"vdm/core/dependencies/mailer"
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/initiate"
	Method = fiber.MethodPost
)

func Route(cfg env.SecurityConfig, mailer mailer.Mailer) *fiberx.Route {
	h := &handler{
		emailVerificationTokenSecret: cfg.EmailVerificationTokenSecret,
		emailVerificationTokenTTL:    cfg.EmailVerificationTokenTTL,
		mailer:                       mailer,
	}
	return fiberx.NewRoute(Method, Path, h.inquireEmailVerification)
}
