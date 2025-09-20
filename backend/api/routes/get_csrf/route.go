package get_csrf

import (
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/csrf-token"
	Method = fiber.MethodGet
)

func Route(cfg env.SecurityConfig) *fiberx.Route {
	handler := &handler{
		csrfTokenSecret: cfg.CsrfTokenSecret,
		csrfTokenTTL:    cfg.CsrfTokenTTL,
	}

	return fiberx.NewRoute(Method, Path, handler.getCsrf)
}
