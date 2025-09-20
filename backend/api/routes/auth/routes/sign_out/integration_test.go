package sign_out

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"vdm/core/dependencies/database"
	"vdm/core/env"
	"vdm/core/fiberx"
	"vdm/core/jwt_utils"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/models"
	"vdm/test_utils"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var testUser = &models.User{ID: uuid.New(), Email: "signout_user0@email.com", EmailVerified: true}

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

	// create a user and one refresh token
	if err = db.Create(testUser).Error; err != nil {
		return
	}
	rft := &models.RefreshToken{UserID: testUser.ID, Expiry: time.Now().Add(10 * time.Minute)}
	if err = db.Create(rft).Error; err != nil {
		return
	}

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

func TestIntegration_SignOut_Success(t *testing.T) {
	c := context.Background()
	container, connector := loadTestData(c, t)
	t.Cleanup(func() { cleanupTestData(c, t, container, connector) })

	app := fiberx.NewApp()

	dummyCfg := env.SecurityConfig{AccessTokenSecret: []byte("dummySecret")}

	Route(connector.GormDB(), dummyCfg).Register(app)

	// Generate access token for the user using the same secret as in route (defaults are OK in tests)
	jwt, err := jwt_utils.GenerateJWT(locals.NewAuthedUser(*testUser), dummyCfg.AccessTokenSecret, time.Now().Add(time.Minute))
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(Method, Path, nil)
	req.AddCookie(&http.Cookie{Name: local_keys.AccessToken, Value: jwt})

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusNoContent, res.StatusCode)

	// Ensure refresh tokens are deleted for this user
	var count int64
	if err := connector.GormDB().Model(&models.RefreshToken{}).Where("user_id = ?", testUser.ID).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), count)
}
