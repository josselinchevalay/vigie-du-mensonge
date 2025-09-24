package create_draft_article

import (
	"vdm/api/middlewares/authorize_authed_user"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/core/env"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	Path   = "/"
	Method = fiber.MethodPost
)

func Group(cfg env.SecurityConfig, db *gorm.DB) *fiberx.Group {
	repo := &repostiroy{db}
	svc := &service{repo}
	handler := &handler{svc}

	group := fiberx.NewGroup(Path)

	group.Add(
		locals_authed_user.Middleware(cfg),
		authorize_authed_user.Middleware(db),
		fiberx.NewRoute(Method, Path, handler.createDraftArticle),
	)

	return group
}
