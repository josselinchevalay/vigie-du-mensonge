package api

import (
	"strings"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/api/middlewares/validate_csrf"
	"vdm/api/routes/auth"
	"vdm/api/routes/email_verification"
	"vdm/api/routes/get_csrf"
	"vdm/api/routes/password_update"
	"vdm/core/dependencies"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

const Prefix = "/api/v1"

func Group(deps *dependencies.Dependencies) *fiberx.Group {
	group := fiberx.NewGroup(Prefix)

	group.Add(
		fiberx.NewMiddleware(helmet.New(helmet.Config{
			XFrameOptions:  "DENY", // stricter than SAMEORIGIN
			ReferrerPolicy: "strict-origin-when-cross-origin",
			// Minimal, valid CSP for API responses (avoid literal "none")
			ContentSecurityPolicy:     "default-src 'none'; frame-ancestors 'none'; base-uri 'none'",
			CrossOriginResourcePolicy: "same-site",
			CrossOriginOpenerPolicy:   "same-origin",
			// Disable COEP for the API to avoid breaking cross-origin fetch/embeds
			CrossOriginEmbedderPolicy: "", // or "unsafe-none" depending on Fiber version
			PermissionPolicy:          "geolocation=(), camera=(), microphone=(), payment=(), usb=()",
		})),
	)

	if deps.Config.ActiveProfile != "prod" {
		group.Add(
			fiberx.NewMiddleware(cors.New(cors.Config{
				AllowOrigins: deps.Config.ClientURL,
				AllowHeaders: strings.Join([]string{
					fiber.HeaderAccessControlAllowOrigin,
					fiber.HeaderOrigin,
					fiber.HeaderContentType,
					fiber.HeaderAccept,
					fiber.HeaderCookie,
					"X-Csrf-Token",
				}, ","),
				AllowCredentials: true,
			})),
		)
	}

	group.Add(
		validate_csrf.Middleware(deps.Config.Security),
		get_csrf.Group(),

		/* /get-csrf has its own rate limiter */
		fiberx.NewMiddleware(limiter.New(limiter.Config{
			Max: 10,
		})),

		auth.Group(deps),
		password_update.Group(deps),

		locals_authed_user.Middleware(deps.Config.Security),

		email_verification.Group(deps),
	)

	return group
}

func Register(router fiber.Router, deps *dependencies.Dependencies) {
	Group(deps).Register(router)
}
