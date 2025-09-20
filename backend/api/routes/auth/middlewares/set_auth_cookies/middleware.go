package set_auth_cookies

import (
	"vdm/core/env"
	"vdm/core/fiberx"
)

func Middleware() *fiberx.Middleware {
	handler := &handler{isProd: env.Config.ActiveProfile == "prod"}

	if handler.isProd {
		handler.sameSite = "Strict"
	} else {
		handler.sameSite = "Lax"
	}

	return fiberx.NewMiddleware(handler.setAuthCookies)
}
