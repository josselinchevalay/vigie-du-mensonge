package process_email_verification

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
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func loadProcessTestData(c context.Context, t *testing.T) (testcontainers.Container, database.Connector, models.User) {
	container, connector := test_utils.NewTestContainerConnector(c, t)
	_db := connector.GormDB()

	user := models.User{
		ID:            uuid.New(),
		Email:         "verify_user@example.com",
		Password:      "irrelevant",
		EmailVerified: false,
	}
	if err := _db.Create(&user).Error; err != nil {
		cleanupTestData(c, t, container, connector)
		t.Fatal(err)
	}
	return container, connector, user
}

func cleanupTestData(c context.Context, t *testing.T, container testcontainers.Container, connector database.Connector) {
	if err := connector.Close(); err != nil {
		t.Logf("failed to close connector: %v", err)
	}
	if err := container.Terminate(c); err != nil {
		t.Logf("failed to terminate container: %v", err)
	}
}

func TestIntegration_EmailVerification_Process_Success(t *testing.T) {
	c := context.Background()
	container, connector, user := loadProcessTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	cfg := env.SecurityConfig{
		EmailVerificationTokenSecret: []byte("secret"),
		EmailVerificationTokenTTL:    1 * time.Minute,
	}

	// Inject authed user middleware
	authed := locals.AuthedUser{ID: user.ID, Email: user.Email}
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("authedUser", authed)
		return c.Next()
	})

	Route(cfg, connector.GormDB()).Register(app)

	// Generate valid token for the authed user
	token, err := jwt_utils.GenerateJWT(authed, cfg.EmailVerificationTokenSecret, time.Now().Add(cfg.EmailVerificationTokenTTL))
	if err != nil {
		t.Fatal(err)
	}

	body, _ := json.Marshal(map[string]string{"token": token})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)

	// Verify DB updated
	var refreshed models.User
	if err := connector.GormDB().First(&refreshed, "id = ?", user.ID).Error; err != nil {
		t.Fatal(err)
	}
	assert.True(t, refreshed.EmailVerified)
}
