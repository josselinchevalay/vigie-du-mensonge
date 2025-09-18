package fiberx

import "github.com/gofiber/fiber/v2"

type Middleware struct {
	handlerFunc fiber.Handler
}

func (m *Middleware) Register(router fiber.Router) {
	router.Use(m.handlerFunc)
}

func NewMiddleware(handlerFunc fiber.Handler) *Middleware {
	return &Middleware{
		handlerFunc: handlerFunc,
	}
}
