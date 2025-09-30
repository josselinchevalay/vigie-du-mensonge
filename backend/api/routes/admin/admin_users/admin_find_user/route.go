package admin_find_user

import (
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

const (
	Path   = "/:" + local_keys.UserTag
	Method = fiber.MethodGet
)
