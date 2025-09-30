package moderator_save_review

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"
	"vdm/core/locals"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type nullRepo struct{}

func (*nullRepo) createReviewAndUpdateArticle(review *models.ArticleReview) error {
	return nil
}

func newAppWithAuthedModeratorAndNullRepo() *fiber.App {
	app := fiberx.NewApp()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("authedUser", locals.AuthedUser{ID: uuid.New()})
		return c.Next()
	})
	h := &handler{repo: &nullRepo{}}
	app.Add(Method, Path, h.saveArticleReviewForModerator)
	return app
}

func TestHandler_ErrBadRequest_ArticleID(t *testing.T) {
	app := newAppWithAuthedModeratorAndNullRepo()
	payload := map[string]any{"decision": string(models.ArticleStatusPublished)}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(Method, "/not-a-uuid/review", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}

func TestHandler_ErrBadRequest_Validate(t *testing.T) {
	app := newAppWithAuthedModeratorAndNullRepo()
	// missing decision
	payload := map[string]any{"notes": "some notes"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(Method, "/"+uuid.New().String()+"/review", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}

func TestHandler_ErrBadRequest_InvalidDecision(t *testing.T) {
	app := newAppWithAuthedModeratorAndNullRepo()
	payload := map[string]any{"decision": "INVALID"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(Method, "/"+uuid.New().String()+"/review", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}

func TestHandler_ErrBadRequest_InvalidNotes(t *testing.T) {
	app := newAppWithAuthedModeratorAndNullRepo()
	payload := map[string]any{"decision": string(models.ArticleStatusChangeRequested), "notes": "too short"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(Method, "/"+uuid.New().String()+"/review", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}
