package api

import (
	"vdm/api/middlewares/cors"
	"vdm/api/routes/auth"
	"vdm/core/dependencies"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
)

const Prefix = "/api/v1"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		cors.Middleware(),
		auth.Group(deps),
	)

	return group
}

func Register(router fiber.Router, deps *dependencies.Dependencies) {
	Group(deps).Register(router)
}
