package inquire_password_update

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Inquire_NoAuthedUser_ReturnsInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := test_utils.NewMockMailer(ctrl)
	m.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	h := &handler{
		passwordUpdateTokenSecret: []byte("secret"),
		passwordUpdateTokenTTL:    time.Minute,
		clientURL:                 "https://example.com",
		mailer:                    m,
	}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.inquirePasswordUpdate)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
}

func TestHandler_Inquire_MailerFails_ReturnsInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	m := test_utils.NewMockMailer(ctrl)
	m.EXPECT().Send(authed.Email, "Mise à jour de votre mot de passe", gomock.Any()).Return(errors.New("send failed")).Times(1)

	h := &handler{
		passwordUpdateTokenSecret: []byte("secret"),
		passwordUpdateTokenTTL:    time.Minute,
		clientURL:                 "https://example.com",
		mailer:                    m,
	}

	app := fiberx.NewApp()
	// inject authed user like the middleware would do
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	app.Add(Method, Path, h.inquirePasswordUpdate)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
}

func TestHandler_Inquire_Success_ReturnsNoContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authed := locals.AuthedUser{ID: uuid.New(), Email: "user@example.com"}

	m := test_utils.NewMockMailer(ctrl)
	m.EXPECT().Send(authed.Email, "Mise à jour de votre mot de passe", gomock.Any()).Return(nil).Times(1)

	h := &handler{
		passwordUpdateTokenSecret: []byte("secret"),
		passwordUpdateTokenTTL:    time.Minute,
		clientURL:                 "https://example.com",
		mailer:                    m,
	}

	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})
	app.Add(Method, Path, h.inquirePasswordUpdate)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}
