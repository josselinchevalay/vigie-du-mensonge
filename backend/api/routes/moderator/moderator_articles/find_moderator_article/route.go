package find_moderator_article

import (
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/:" + local_keys.ArticleReference
	Method = fiber.MethodGet
)
