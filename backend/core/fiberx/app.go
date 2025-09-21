package fiberx

import (
	"errors"
	"vdm/core/logger"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder: sonic.ConfigFastest.Marshal,
		JSONDecoder: sonic.ConfigFastest.Unmarshal,
		BodyLimit:   2 * 1024 * 1024,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error("error handling request",
				logger.Any("path", c.Path()),
				logger.Err(err))

			var code int

			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				code = fiberErr.Code
			} else {
				code = fiber.StatusInternalServerError
			}

			return c.SendStatus(code)
		},
	})
}
