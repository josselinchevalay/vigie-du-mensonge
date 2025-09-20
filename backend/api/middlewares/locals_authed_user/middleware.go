package locals_authed_user

import (
	"vdm/core/env"
	"vdm/core/fiberx"
)

func Middleware(cfg env.SecurityConfig) *fiberx.Middleware {
	handler := &handler{cfg.AccessTokenSecret}

	return fiberx.NewMiddleware(handler.localsAuthedUser)
}
