package save_redactor_article

import (
	"vdm/core/locals"
	"vdm/core/models"
	"vdm/core/validation"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	saveArticleForRedactor(c *fiber.Ctx) error
}

type handler struct {
	svc Service
}

func (h *handler) saveArticleForRedactor(c *fiber.Ctx) error {
	authedUser, ok := c.Locals("authedUser").(locals.AuthedUser)
	if !ok {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "can't locals authed user"}
	}

	var reqDTO RequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid request body"}
	}
	if err := validation.Validate(reqDTO); err != nil {
		return err
	}

	article, err := reqDTO.toArticle(authedUser.ID)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	publish := c.Query("publish") == "true"

	if publish {
		if err = validateForPublication(article); err != nil {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
		}
	}

	if err = h.svc.saveArticleForRedactor(publish, article); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func validateForPublication(article models.Article) error {
	if len(article.Body) < 200 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "article body is too short"}
	}
	if len(article.Body) > 2000 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "article body is too long"}
	}

	if len(article.Tags) < 1 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "article must have at least one tag"}
	}

	if len(article.Sources) < 1 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "article must have at least one source"}
	}

	return nil
}
