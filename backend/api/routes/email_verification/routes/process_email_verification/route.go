package process_email_verification

import (
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/process"
	Method = fiber.MethodPost
)

func Route(cfg env.SecurityConfig, db *gorm.DB) *fiberx.Route {
	repo := &repository{db}
	svc := &service{
		repo:        repo,
		tokenSecret: cfg.EmailTokenSecret,
	}
	handler := &handler{svc}
	return fiberx.NewRoute(Method, Path, handler.processEmailVerification)
}
