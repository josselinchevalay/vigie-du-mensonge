package refresh

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type nullService struct{}

func (nullService) refresh(rftID uuid.UUID) (models.User, locals.AccessToken, locals.RefreshToken, error) {
	return models.User{}, locals.AccessToken{}, locals.RefreshToken{}, nil
}

func TestHandler_ErrBadRequest_MissingCookie(t *testing.T) {
	h := &handler{svc: nullService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.refresh)

	req := httptest.NewRequest(Method, Path, nil)
	// No cookie set

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestHandler_ErrBadRequest_InvalidCookie(t *testing.T) {
	h := &handler{svc: nullService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.refresh)

	req := httptest.NewRequest(Method, Path, nil)
	req.AddCookie(&http.Cookie{Name: local_keys.RefreshToken, Value: "not-a-uuid"})

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
