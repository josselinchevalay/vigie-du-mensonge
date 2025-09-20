package locals_authed_user

import (
	"vdm/core/env"
	"vdm/core/fiberx"
)

func Middleware() *fiberx.Middleware {
	handler := &handler{env.Config.Security.AccessTokenSecret}

	return fiberx.NewMiddleware(handler.localsAuthedUser)
}
