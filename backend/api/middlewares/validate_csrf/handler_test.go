package validate_csrf

import (
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/fiberx"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware_ValidToken_AllowsRequest(t *testing.T) {
	secret := []byte("csrf-secret")
	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	app := fiberx.NewApp()
	// Inject authed user before CSRF middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	Middleware(secret).Register(app)
	app.Post("/protected", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	req := httptest.NewRequest(fiber.MethodPost, "/protected", nil)
	if token, err := jwt_utils.GenerateJWT(authed, secret, time.Now().Add(time.Minute)); err != nil {
		t.Fatal(err)
	} else {
		req.Header.Set("csrf-token", token)
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}

func TestMiddleware_MissingToken_ReturnsUnauthorized(t *testing.T) {
	secret := []byte("csrf-secret")
	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	Middleware(secret).Register(app)
	app.Post("/protected", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	req := httptest.NewRequest(fiber.MethodPost, "/protected", nil)
	// No csrf-token header

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
}

func TestMiddleware_InvalidSignature_ReturnsUnauthorized(t *testing.T) {
	secret := []byte("csrf-secret")
	badSecret := []byte("other-secret")
	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	Middleware(secret).Register(app)
	app.Post("/protected", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	req := httptest.NewRequest(fiber.MethodPost, "/protected", nil)
	if token, err := jwt_utils.GenerateJWT(authed, badSecret, time.Now().Add(time.Minute)); err != nil {
		t.Fatal(err)
	} else {
		req.Header.Set("csrf-token", token)
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
}

func TestMiddleware_MismatchedUser_ReturnsForbidden(t *testing.T) {
	secret := []byte("csrf-secret")
	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}
	other := locals.AuthedUser{ID: uuid.New(), Email: "other@example.com"}

	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	Middleware(secret).Register(app)
	app.Post("/protected", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	req := httptest.NewRequest(fiber.MethodPost, "/protected", nil)
	if token, err := jwt_utils.GenerateJWT(other, secret, time.Now().Add(time.Minute)); err != nil {
		t.Fatal(err)
	} else {
		req.Header.Set("csrf-token", token)
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusForbidden, res.StatusCode)
}
