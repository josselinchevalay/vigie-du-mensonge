package sign_up

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/dependencies/database"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var testUser = &models.User{Email: "user0@email.com"}

func loadTestData(c context.Context, t *testing.T) (container testcontainers.Container, connector database.Connector) {
	container, connector = test_utils.NewTestContainerConnector(c, t)

	db := connector.GormDB()

	var err error

	defer func(c context.Context, container testcontainers.Container, connector database.Connector) {
		if err != nil {
			cleanupTestData(c, t, container, connector)
			t.Fatal(err)
		}
	}(c, container, connector)

	err = db.Create(testUser).Error

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
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret"), AccessTokenTTL: 1 * time.Minute, RefreshTokenTTL: 1 * time.Minute}

	Route(connector.GormDB(), dummyCfg).Register(app)

	reqDTO := RequestDTO{
		Email:    "user1@email.com",
		Password: "Test123!",
	}

	reqBody, err := json.Marshal(reqDTO)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(Method, Path, bytes.NewReader(reqBody))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusCreated, res.StatusCode)
}

func TestIntegration_ErrConflict(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret"), AccessTokenTTL: 1 * time.Minute, RefreshTokenTTL: 1 * time.Minute}

	Route(connector.GormDB(), dummyCfg).Register(app)

	reqDTO := RequestDTO{
		Email:    testUser.Email,
		Password: "Test123!",
	}

	reqBody, err := json.Marshal(reqDTO)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(Method, Path, bytes.NewReader(reqBody))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusConflict, res.StatusCode)
}
