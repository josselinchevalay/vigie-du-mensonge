package fiberx

import "github.com/gofiber/fiber/v2"

type Route struct {
	method, path string
	handlerFunc  fiber.Handler
}

func (r *Route) Register(router fiber.Router) {
	router.Add(r.method, r.path, r.handlerFunc)
}

func NewRoute(method, path string, handlerFunc fiber.Handler) *Route {
	return &Route{
		method:      method,
		path:        path,
		handlerFunc: handlerFunc,
	}
}
