package inquire_email_verification

import (
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type nullService struct{}

func (*nullService) sendEmailAndCreateToken(authedUser locals.AuthedUser) error {
	return nil
}

func TestHandler_ErrConflict(t *testing.T) {
	handler := &handler{&nullService{}}

	app := fiberx.NewApp()

	app.Add(Method, Path, func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, locals.AuthedUser{EmailVerified: true})
		return c.Next()
	}, handler.inquireEmailVerification)

	req := httptest.NewRequest(Method, Path, nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusConflict, res.StatusCode)
}
