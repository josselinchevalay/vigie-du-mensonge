package authorize_authed_user

import (
	"vdm/core/fiberx"

	"gorm.io/gorm"
)

func Middleware(db *gorm.DB) *fiberx.Middleware {
	repo := &repository{db}
	svc := &service{repo}
	handler := &handler{svc}

	return fiberx.NewMiddleware(handler.authorizedAuthedUser)
}
