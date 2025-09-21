package dev_cors

import (
	"strings"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// for development purposes only
// not required for production since client is hosted on the same domain
func Middleware(allowOrigins string) *fiberx.Middleware {
	return fiberx.NewMiddleware(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     allowHeaders(),
		AllowCredentials: true,
	}))
}

func allowHeaders() string {
	headers := []string{
		fiber.HeaderContentType,
		"X-Csrf-Token",
	}

	return strings.Join(headers, ",")
}
