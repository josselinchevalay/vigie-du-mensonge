package moderator_claim_article

import (
	"fmt"
	"vdm/core/locals"
	"vdm/core/locals/local_keys"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	claimArticleForModerator(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) claimArticleForModerator(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articleID, err := uuid.Parse(c.Params(local_keys.ArticleID))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article id"}
	}

	if err := h.repo.updateArticleModerator(authedUser.ID, articleID); err != nil {
		return fmt.Errorf("failed to update article moderator: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
