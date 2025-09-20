package locals_authed_user

import (
	"vdm/core/jwt_utils"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	localsAuthedUser(c *fiber.Ctx) error
}

type handler struct {
	accessTokenSecret []byte
}

func (h *handler) localsAuthedUser(c *fiber.Ctx) error {
	authedUser, err := jwt_utils.ParseJWT(c.Cookies(local_keys.AccessToken), h.accessTokenSecret)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}

	c.Locals(local_keys.AuthedUser, authedUser)

	return c.Next()
}
