package redactor_save_article

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type nullService struct{}

func (*nullService) saveArticleForRedactor(publish bool, newArticle models.Article) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func newAppWithAuthedUserAndNullSvc() *fiber.App {
	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("authedUser", locals.AuthedUser{ID: uuid.New()})
		return c.Next()
	})
	h := &handler{svc: &nullService{}}
	app.Add(Method, Path, h.saveArticleForRedactor)
	return app
}

func TestHandler_ErrBadRequest_ValidateErr(t *testing.T) {
	app := newAppWithAuthedUserAndNullSvc()

	// invalid: title too short (min 20) and missing category
	payload := map[string]any{
		"title":     "short title",
		"eventDate": time.Now(),
	}
	b, _ := json.Marshal(payload)

	req := httptest.NewRequest(Method, Path, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestHandler_ErrBadRequest_MappingErr(t *testing.T) {
	app := newAppWithAuthedUserAndNullSvc()

	// invalid: duplicate tag to trigger mapping error
	payload := map[string]any{
		"title":         "This is a sufficiently long article title",
		"eventDate":     time.Now(),
		"category":      string(models.ArticleCategoryLie),
		"tags":          []string{"dup", "dup"},
		"politicianIds": []string{},
		"sources":       []string{},
	}
	b, _ := json.Marshal(payload)

	req := httptest.NewRequest(Method, Path, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestHandler_ErrBadRequest_ValidateForPublicationErr(t *testing.T) {
	app := newAppWithAuthedUserAndNullSvc()

	// publish=true but body too short and no tags/sources
	payload := map[string]any{
		"title":         "This is a sufficiently long article title",
		"eventDate":     time.Now(),
		"category":      string(models.ArticleCategoryLie),
		"body":          "too short",
		"tags":          []string{},
		"politicianIds": []string{},
		"sources":       []string{},
	}
	b, _ := json.Marshal(payload)

	req := httptest.NewRequest(Method, Path+"?publish=true", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
