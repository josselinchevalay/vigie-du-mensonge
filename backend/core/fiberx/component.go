package fiberx

import "github.com/gofiber/fiber/v2"

type Component interface {
	Register(router fiber.Router)
}
