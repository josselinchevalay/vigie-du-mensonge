package fiberx

import "github.com/gofiber/fiber/v2"

type Group struct {
	prefix     string
	components []Component
}

func (g *Group) Register(router fiber.Router) {
	router = router.Group(g.prefix)

	for _, component := range g.components {
		component.Register(router)
	}
}

func NewGroup(prefix string) *Group {
	return &Group{prefix: prefix}
}

func (g *Group) Add(components ...Component) {
	g.components = append(g.components, components...)
}
