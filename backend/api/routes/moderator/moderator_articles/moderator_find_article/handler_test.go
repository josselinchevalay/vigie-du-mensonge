package moderator_find_article

import (
	"net/http/httptest"
	"testing"
	"vdm/core/fiberx"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type nullRepo struct{}

func (*nullRepo) findArticlesByReference(reference uuid.UUID) ([]models.Article, error) {
	return nil, nil
}

func TestHandler_ErrBadRequest(t *testing.T) {
	app := fiberx.NewApp()
	h := &handler{repo: &nullRepo{}}
	app.Add(Method, Path, h.findArticlesByReferenceForModerator)

	req := httptest.NewRequest(Method, "/not-a-uuid", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}
