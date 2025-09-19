package set_auth_cookies

import (
	"vdm/core/env"
	"vdm/core/fiberx"
)

func Middleware() *fiberx.Middleware {
	handler := &handler{env.Config.ActiveProfile == "prod"}
	return fiberx.NewMiddleware(handler.setAuthCookies)
}
