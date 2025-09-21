package inquire_password_update

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
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func loadTestData(c context.Context, t *testing.T) (testcontainers.Container, database.Connector, models.User) {
	container, connector := test_utils.NewTestContainerConnector(c, t)
	_db := connector.GormDB()

	user := models.User{
		ID:            uuid.New(),
		Email:         "pwd_inquire_user@example.com",
		Password:      "$2a$10$initialhashedpasswordplaceholder0123456789012345", // placeholder only
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

func TestIntegration_Success(t *testing.T) {
	c := context.Background()
	container, connector, user := loadTestData(c, t)
	t.Cleanup(func() { cleanUpTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	cfg := env.SecurityConfig{
		PasswordUpdateTokenSecret: []byte("secret"),
		PasswordUpdateTokenTTL:    1 * time.Minute,
	}
	clientURL := "https://client.example"

	mockCtrl := gomock.NewController(t)
	mailer := test_utils.NewMockMailer(mockCtrl)

	Route(cfg, clientURL, connector.GormDB(), mailer).Register(app)

	body, _ := json.Marshal(RequestDTO{Email: user.Email})
	req := httptest.NewRequest(Method, Path, bytes.NewReader(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	mailer.EXPECT().Send(user.Email, gomock.Any(), gomock.Any()).Return(nil)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)
}
