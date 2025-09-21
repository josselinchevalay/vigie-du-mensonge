package validate_csrf

import (
	"time"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func Middleware(cfg env.SecurityConfig) *fiberx.Middleware {
	return fiberx.NewMiddleware(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     cfg.CsrfCookieName,
		CookieSameSite: cfg.CookieSameSite,
		CookieSecure:   cfg.CookieSecure,
		CookieHTTPOnly: true,
		SingleUseToken: true,
		CookiePath:     "/",
		Expiration:     1 * time.Minute,
		KeyGenerator:   utils.UUIDv4,
		ContextKey:     local_keys.CsrfToken,
	}))
}
