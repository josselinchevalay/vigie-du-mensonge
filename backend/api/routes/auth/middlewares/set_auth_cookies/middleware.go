package set_auth_cookies

import (
	"vdm/core/fiberx"
)

func Middleware(isProd bool) *fiberx.Middleware {
	handler := &handler{isProd: isProd}

	if handler.isProd {
		handler.sameSite = "Strict"
	} else {
		handler.sameSite = "Lax"
	}

	return fiberx.NewMiddleware(handler.setAuthCookies)
}
