package cors

import (
	"strings"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Middleware(allowOrigins string) *fiberx.Middleware {
	return fiberx.NewMiddleware(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     allowHeaders(),
		AllowCredentials: true,
	}))
}

func allowHeaders() string {
	headers := []string{
		fiber.HeaderAccessControlAllowOrigin,
		fiber.HeaderOrigin,
		fiber.HeaderContentType,
		fiber.HeaderAccept,
		fiber.HeaderCookie,
		"csrf-token",
	}

	return strings.Join(headers, ",")
}
