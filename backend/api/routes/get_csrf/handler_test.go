package get_csrf

import (
	"encoding/json"
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

type csrfResponse struct {
	CsrfToken string `json:"csrfToken"`
}

func TestHandler_GetCsrf_NoAuthedUser_ReturnsInternalServerError(t *testing.T) {
	h := &handler{csrfTokenSecret: []byte("secret"), csrfTokenTTL: time.Minute}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.getCsrf)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
}

func TestHandler_GetCsrf_Success_ReturnsToken(t *testing.T) {
	secret := []byte("secret")
	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	h := &handler{csrfTokenSecret: secret, csrfTokenTTL: time.Minute}

	app := fiberx.NewApp()
	// Inject authed user like auth middleware would do
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	app.Add(Method, Path, h.getCsrf)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	var body csrfResponse
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body.CsrfToken == "" {
		t.Fatalf("expected non-empty csrfToken")
	}

	// Validate the token can be parsed and matches authed user
	parsedUser, err := jwt_utils.ParseJWT(body.CsrfToken, secret)
	if err != nil {
		t.Fatalf("failed to parse csrf token: %v", err)
	}
	assert.Equal(t, authed, parsedUser)
}
