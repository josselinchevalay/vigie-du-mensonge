package sign_up

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type nullService struct{}

func (nullService) signUp(req RequestDTO) (locals.AccessToken, locals.RefreshToken, error) {
	return locals.AccessToken{}, locals.RefreshToken{}, nil
}

func TestHandler_ErrBadRequest(t *testing.T) {
	h := &handler{svc: nullService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.signUp)

	req := httptest.NewRequest(Method, Path, nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestHandler_ErrBadRequest2(t *testing.T) {
	h := &handler{svc: nullService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.signUp)

	reqBody, err := json.Marshal(RequestDTO{})
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(Method, Path, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
