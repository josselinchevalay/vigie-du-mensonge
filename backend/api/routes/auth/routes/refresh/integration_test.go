package refresh

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/dependencies/database"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/locals/local_keys"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var testRoles = []*models.Role{
	{Name: "ADMIN"},
	{Name: "MODERATOR"},
}
var testUser = &models.User{ID: uuid.New(), Email: "refresh_user0@email.com", EmailVerified: true, Roles: testRoles}

var validRft = &models.RefreshToken{UserID: testUser.ID, Expiry: time.Now().Add(1 * time.Minute)}
var expiredRft = &models.RefreshToken{UserID: testUser.ID, Expiry: time.Now().Add(-1 * time.Minute)}

func loadTestData(c context.Context, t *testing.T) (container testcontainers.Container, connector database.Connector) {
	container, connector = test_utils.NewTestContainerConnector(c, t)

	db := connector.GormDB()

	var err error

	defer func() {
		if err != nil {
			cleanupTestData(c, t, container, connector)
			t.Fatal(err)
		}
	}()

	if err = db.Create(testRoles).Error; err != nil {
		return
	}

	if err = db.Create(testUser).Error; err != nil {
		return
	}

	err = db.Create([]*models.RefreshToken{validRft, expiredRft}).Error

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

func TestIntegration_Refresh_Success(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret"), AccessTokenTTL: 1 * time.Minute, RefreshTokenTTL: 1 * time.Minute}

	Route(connector.GormDB(), dummyCfg).Register(app)

	req := httptest.NewRequest(Method, Path, nil)
	req.AddCookie(&http.Cookie{Name: local_keys.RefreshToken, Value: validRft.ID.String()})
	// no body required

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resDTO RefreshResponse
	if err = json.Unmarshal(resBody, &resDTO); err != nil {
		t.Fatal(err)
	}

	assert.True(t, resDTO.EmailVerified)
	assert.Equal(t, testUser.MapRoles(), resDTO.Roles)
}

func TestIntegration_Refresh_ErrUnauthorized(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret"), AccessTokenTTL: 1 * time.Minute, RefreshTokenTTL: 1 * time.Minute}

	Route(connector.GormDB(), dummyCfg).Register(app)

	req := httptest.NewRequest(Method, Path, nil)
	req.AddCookie(&http.Cookie{Name: local_keys.RefreshToken, Value: expiredRft.ID.String()})
	// no body required

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
}
