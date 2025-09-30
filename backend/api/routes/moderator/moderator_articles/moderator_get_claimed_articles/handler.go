package moderator_get_claimed_articles

import (
	"fmt"
	"vdm/core/dto/response_dto"
	"vdm/core/locals"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	getClaimedArticlesForModerator(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) getClaimedArticlesForModerator(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articles, err := h.repo.getArticlesByModeratorID(authedUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get moderator articles: %v", err)
	}

	resDTO := make([]response_dto.Article, len(articles))
	for i := range articles {
		resDTO[i] = response_dto.NewArticle(articles[i])
	}

	return c.Status(fiber.StatusOK).JSON(resDTO)
}
