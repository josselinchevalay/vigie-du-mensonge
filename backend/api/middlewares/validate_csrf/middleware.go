package validate_csrf

import (
	"time"
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func Middleware(isProd bool) *fiberx.Middleware {
	var sameSite string
	if isProd {
		sameSite = "Strict"
	} else {
		sameSite = "Lax"
	}

	return fiberx.NewMiddleware(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "__Host-csrf_",
		CookieSameSite: sameSite,
		CookieSecure:   isProd,
		CookieHTTPOnly: true,
		SingleUseToken: true,
		CookiePath:     "/",
		Expiration:     1 * time.Minute,
		KeyGenerator:   utils.UUIDv4,
		ContextKey:     local_keys.CsrfToken,
	}))
}
