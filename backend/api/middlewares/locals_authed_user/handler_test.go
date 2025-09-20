package locals_authed_user

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Success(t *testing.T) {
	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret")}

	Middleware(dummyCfg).Register(app)

	input := locals.AuthedUser{ID: uuid.New(), Email: "test@email.com"}

	app.Get("/", func(c *fiber.Ctx) error {
		output, ok := c.Locals(local_keys.AuthedUser, input).(locals.AuthedUser)
		if !ok {
			return fiber.ErrInternalServerError
		}

		assert.Equal(t, input, output)

		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)

	if jwt, err := jwt_utils.GenerateJWT(input, dummyCfg.AccessTokenSecret, time.Now().Add(time.Minute)); err != nil {
		t.Fatal(err)
	} else {
		req.AddCookie(&http.Cookie{Name: local_keys.AccessToken, Value: jwt})
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}

func TestHandler_ErrUnauthorized(t *testing.T) {
	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret")}

	Middleware(dummyCfg).Register(app)

	input := locals.AuthedUser{ID: uuid.New(), Email: "test@email.com"}

	app.Get("/", func(c *fiber.Ctx) error {
		output, ok := c.Locals(local_keys.AuthedUser, input).(locals.AuthedUser)
		if !ok {
			return fiber.ErrInternalServerError
		}

		assert.Equal(t, input, output)

		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)

	if jwt, err := jwt_utils.GenerateJWT(input, dummyCfg.AccessTokenSecret, time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	} else {
		req.AddCookie(&http.Cookie{Name: local_keys.AccessToken, Value: jwt})
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
}
