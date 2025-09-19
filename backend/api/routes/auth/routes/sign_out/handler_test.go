package sign_out

import (
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type nullService struct{}

func (nullService) signOut(accessToken string) {}

func TestHandler_SignOut_NoCookie_ReturnsNoContent(t *testing.T) {
	h := &handler{svc: nullService{}}

	app := fiberx.NewApp()
	app.Add(Method, Path, h.signOut)

	req := httptest.NewRequest(Method, Path, nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}
