package save_moderator_review

import (
	"vdm/core/locals"
	"vdm/core/locals/local_keys"
	"vdm/core/models"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler interface {
	saveArticleReviewForModerator(c *fiber.Ctx) error
}

type handler struct {
	repo Repository
}

func (h *handler) saveArticleReviewForModerator(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	articleID, err := uuid.Parse(c.Params(local_keys.ArticleID))
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid article id"}
	}

	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(reqDTO); err != nil {
		return err
	}

	if !reqDTO.Decision.Valid() {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid decision"}
	}

	// optional notes only for published articles
	if reqDTO.Decision != models.ArticleStatusPublished && len(reqDTO.Notes) < 30 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid notes"}
	}

	review := &models.ArticleReview{
		ArticleID:   articleID,
		ModeratorID: authedUser.ID,
		Decision:    reqDTO.Decision,
	}

	if len(reqDTO.Notes) > 0 {
		review.Notes = reqDTO.Notes
	}

	if err := h.repo.createReviewAndUpdateArticle(review); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
