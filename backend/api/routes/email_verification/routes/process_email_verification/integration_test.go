package process_email_verification

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/api/middlewares/locals_authed_user"
	"vdm/core/dependencies/database"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/hmac_utils"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

type testData struct {
	user         *models.User
	validToken   uuid.UUID
	expiredToken uuid.UUID
	cfg          env.SecurityConfig
}

func loadTestData(c context.Context, t *testing.T) (container testcontainers.Container, connector database.Connector, data testData) {
	container, connector = test_utils.NewTestContainerConnector(c, t)
	db := connector.GormDB()

	data.user = &models.User{
		Email:         "verify_user@example.com",
		EmailVerified: false,
	}

	var err error

	defer func(c context.Context, container testcontainers.Container, connector database.Connector) {
		if err != nil {
			cleanupTestData(c, t, container, connector)
			t.Fatal(err)
		}
	}(c, container, connector)

	if err = db.Create(data.user).Error; err != nil {
		return
	}

	data.cfg = env.SecurityConfig{
		AccessTokenSecret: []byte("access"),
		AccessTokenTTL:    1 * time.Minute,
		AccessCookieName:  "access",
		EmailTokenSecret:  []byte("email"),
		EmailTokenTTL:     1 * time.Minute,
	}

	data.validToken = uuid.New()
	validUsrTok := models.UserToken{UserID: data.user.ID, Hash: hmac_utils.HashUUID(data.validToken, data.cfg.EmailTokenSecret),
		Expiry: time.Now().Add(data.cfg.EmailTokenTTL), Category: models.UserTokenCategoryEmail}

	data.expiredToken = uuid.New()
	expiredUsrTok := models.UserToken{UserID: data.user.ID, Hash: hmac_utils.HashUUID(data.expiredToken, data.cfg.EmailTokenSecret),
		Expiry: time.Now().Add(-1 * time.Minute), Category: models.UserTokenCategoryEmail}

	err = db.Create([]models.UserToken{validUsrTok, expiredUsrTok}).Error

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

	app := fiberx.NewApp()

	locals_authed_user.Middleware(data.cfg).Register(app)
	Route(data.cfg, connector.GormDB()).Register(app)

	body, _ := json.Marshal(RequestDTO{Token: data.validToken})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	authed := locals.AuthedUser{ID: data.user.ID, Email: data.user.Email, EmailVerified: data.user.EmailVerified}
	if jwt, err := jwt_utils.GenerateJWT(authed, data.cfg.AccessTokenSecret, time.Now().Add(data.cfg.AccessTokenTTL)); err != nil {
		t.Fatal(err)
	} else {
		req.AddCookie(&http.Cookie{Name: data.cfg.AccessCookieName, Value: jwt})
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)

	// Verify DB updated
	var verified models.User
	if err := connector.GormDB().First(&verified, "id = ?", data.user.ID).Error; err != nil {
		t.Fatal(err)
	}
	assert.True(t, verified.EmailVerified)

	var tokenCount int
	if err := connector.GormDB().Model(&models.UserToken{}).
		Where("user_id = ?", data.user.ID).
		Select("count(*)").
		Find(&tokenCount).Error; err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, tokenCount)
}

func TestIntegration_ErrGone(t *testing.T) {
	c := context.Background()
	container, connector, data := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	locals_authed_user.Middleware(data.cfg).Register(app)
	Route(data.cfg, connector.GormDB()).Register(app)

	body, _ := json.Marshal(RequestDTO{Token: data.expiredToken})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	authed := locals.AuthedUser{ID: data.user.ID, Email: data.user.Email, EmailVerified: data.user.EmailVerified}
	if jwt, err := jwt_utils.GenerateJWT(authed, data.cfg.AccessTokenSecret, time.Now().Add(data.cfg.AccessTokenTTL)); err != nil {
		t.Fatal(err)
	} else {
		req.AddCookie(&http.Cookie{Name: data.cfg.AccessCookieName, Value: jwt})
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusGone, res.StatusCode)
}
