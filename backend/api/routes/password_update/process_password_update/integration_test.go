package process_password_update

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
	"vdm/core/locals/local_keys"
	"vdm/core/models"
	"vdm/test_utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"golang.org/x/crypto/bcrypt"
)

func loadTestData(c context.Context, t *testing.T) (testcontainers.Container, database.Connector, models.User) {
	container, connector := test_utils.NewTestContainerConnector(c, t)
	_db := connector.GormDB()

	user := models.User{
		ID:            uuid.New(),
		Email:         "pwd_update_user@example.com",
		Password:      "$2a$10$initialhashedpasswordplaceholder0123456789012345", // placeholder, will be overwritten
		EmailVerified: true,
	}
	if err := _db.Create(&user).Error; err != nil {
		cleanUpTestData(c, t, container, connector)
		t.Fatal(err)
	}
	return container, connector, user
}

func cleanUpTestData(c context.Context, t *testing.T, container testcontainers.Container, connector database.Connector) {
	if err := connector.Close(); err != nil {
		t.Logf("failed to close connector: %v", err)
	}
	if err := container.Terminate(c); err != nil {
		t.Logf("failed to terminate container: %v", err)
	}
}

func TestIntegration_PasswordUpdate_Process_Success(t *testing.T) {
	c := context.Background()
	container, connector, user := loadTestData(c, t)
	t.Cleanup(func() { cleanUpTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	cfg := env.SecurityConfig{
		PasswordUpdateTokenSecret: []byte("secret"),
		PasswordUpdateTokenTTL:    1 * time.Minute,
	}

	// Inject authed user into locals as middleware would
	authed := locals.AuthedUser{ID: user.ID, Email: user.Email}
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(local_keys.AuthedUser, authed)
		return c.Next()
	})

	Route(connector.GormDB(), cfg).Register(app)

	// Generate valid password update token for authed user
	token, err := jwt_utils.GenerateJWT(authed, cfg.PasswordUpdateTokenSecret, time.Now().Add(cfg.PasswordUpdateTokenTTL))
	if err != nil {
		t.Fatal(err)
	}

	newPwd := "NewPass123!"
	body, _ := json.Marshal(RequestDTO{Token: token, NewPassword: newPwd})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)

	// Verify DB password updated and matches new password
	var refreshed models.User
	if err := connector.GormDB().First(&refreshed, "id = ?", user.ID).Error; err != nil {
		t.Fatal(err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(refreshed.Password), []byte(newPwd)); err != nil {
		t.Fatalf("expected password to be updated and match new password, compare error: %v", err)
	}
}
