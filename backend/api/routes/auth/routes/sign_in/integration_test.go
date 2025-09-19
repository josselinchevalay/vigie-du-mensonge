package sign_in

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"vdm/core/dependencies/database"
	"vdm/core/fiberx"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"golang.org/x/crypto/bcrypt"
)

var testRoles = []*models.Role{
	{Name: "ADMIN"},
	{Name: "MODERATOR"},
}
var testUser = &models.User{Email: "signin_user0@email.com", Roles: testRoles}

func loadTestData(c context.Context, t *testing.T) (container testcontainers.Container, connector database.Connector) {
	container, connector = database.NewTestContainerConnector(c, t)

	db := connector.GormDB()

	var err error

	defer func() {
		if err != nil {
			cleanupTestData(c, t, container, connector)
			t.Fatal(err)
		}
	}()

	// create a user with known password
	pwd, _ := bcrypt.GenerateFromPassword([]byte("Test123!"), bcrypt.DefaultCost)
	testUser.Password = string(pwd)
	testUser.EmailVerified = true

	if err = db.Create(&testRoles).Error; err != nil {
		t.Fatal(err)
	}

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

func TestIntegration_SignIn_Success(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()
	Route(connector.GormDB()).Register(app)

	reqDTO := SignInRequest{
		Email:    testUser.Email,
		Password: "Test123!",
	}
	b, _ := json.Marshal(reqDTO)

	req := httptest.NewRequest(Method, Path, bytes.NewReader(b))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	// parse body and assert roles length
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var dto SignInResponse
	if err := json.Unmarshal(body, &dto); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(dto.Roles))
}

func TestIntegration_SignIn_Unauthorized(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()
	Route(connector.GormDB()).Register(app)

	reqDTO := SignInRequest{
		Email:    testUser.Email,
		Password: "WrongPassword" + time.Now().String(),
	}
	b, _ := json.Marshal(reqDTO)

	req := httptest.NewRequest(Method, Path, bytes.NewReader(b))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
}

func TestIntegration_SignIn_NotFound(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()
	Route(connector.GormDB()).Register(app)

	reqDTO := SignInRequest{
		Email:    "unknown_" + strconv.FormatInt(time.Now().UnixNano(), 10) + "@email.com",
		Password: "SomePassword1!",
	}
	b, _ := json.Marshal(reqDTO)

	req := httptest.NewRequest(Method, Path, bytes.NewReader(b))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
}
