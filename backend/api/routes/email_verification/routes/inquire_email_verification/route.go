package inquire_email_verification

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

func Route(cfg env.SecurityConfig, db *gorm.DB, clientURL string, mailer mailer.Mailer) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:        repo,
		tokenSecret: cfg.EmailTokenSecret,
		tokenTTL:    cfg.EmailTokenTTL,
		mailer:      mailer,
		clientURL:   clientURL,
	}
	h := &handler{svc}
	return fiberx.NewRoute(Method, Path, h.inquireEmailVerification)
}
