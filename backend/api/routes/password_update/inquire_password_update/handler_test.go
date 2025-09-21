package inquire_password_update

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	retErr     error
	calledWith string
}

func (m *mockService) inquirePasswordUpdate(email string) error {
	m.calledWith = email
	return m.retErr
}

func TestHandler_Inquire_InvalidBody_ReturnsBadRequest(t *testing.T) {
	h := &handler{svc: &mockService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.inquirePasswordUpdate)

	req := httptest.NewRequest(Method, Path, bytes.NewBufferString("not-json"))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestHandler_Inquire_Success_ReturnsNoContent(t *testing.T) {
	m := &mockService{}
	h := &handler{svc: m}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.inquirePasswordUpdate)

	email := "user@example.com"
	body, _ := json.Marshal(RequestDTO{Email: email})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
	assert.Equal(t, email, m.calledWith)
}
