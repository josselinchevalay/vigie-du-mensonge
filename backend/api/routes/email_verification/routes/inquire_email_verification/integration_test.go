package inquire_email_verification

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/core/dependencies/database"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

type testData struct {
	user *models.User
}

func loadTestData(c context.Context, t *testing.T) (container testcontainers.Container, connector database.Connector, data testData) {
	container, connector = test_utils.NewTestContainerConnector(c, t)

	db := connector.GormDB()

	var err error

	defer func(c context.Context, container testcontainers.Container, connector database.Connector) {
		if err != nil {
			cleanupTestData(c, t, container, connector)
			t.Fatal(err)
		}
	}(c, container, connector)

	data.user = &models.User{EmailVerified: false, Email: "verified@user.com"}

	err = db.Create(data.user).Error

	return
}

func cleanupTestData(c context.Context, t *testing.T, container testcontainers.Container, connector database.Connector) {
	if err := connector.Close(); err != nil {
		t.Logf("failed to close connector: %v", err)
	}

	if err := container.Terminate(c); err != nil {
		t.Logf("failed to terminate container: %v", err)
	}
}

func TestIntegration_Success(t *testing.T) {
	c := context.Background()
	container, connector, data := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	db := connector.GormDB()

	mockCtrl := gomock.NewController(t)
	mailer := test_utils.NewMockMailer(mockCtrl)

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{
		AccessTokenSecret: []byte("access"),
		AccessTokenTTL:    1 * time.Minute,
		EmailTokenSecret:  []byte("email"),
		EmailTokenTTL:     1 * time.Minute,
	}
	clientURL := "https://client.example"

	locals_authed_user.Middleware(dummyCfg).Register(app)
	Route(dummyCfg, db, clientURL, mailer).Register(app)

	req := httptest.NewRequest(Method, Path, nil)
	authedUser := locals.AuthedUser{ID: data.user.ID, Email: data.user.Email, EmailVerified: data.user.EmailVerified}

	if jwt, err := jwt_utils.GenerateJWT(authedUser, dummyCfg.AccessTokenSecret, time.Now().Add(dummyCfg.AccessTokenTTL)); err != nil {
		t.Fatal(err)
	} else {
		req.AddCookie(&http.Cookie{Name: dummyCfg.AccessCookieName, Value: jwt})
	}

	mailer.EXPECT().Send(authedUser.Email, gomock.Any(), gomock.Any()).Return(nil).Times(1)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}
