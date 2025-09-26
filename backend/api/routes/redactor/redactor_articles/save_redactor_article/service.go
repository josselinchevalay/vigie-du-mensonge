package save_redactor_article

import (
	"fmt"
	"vdm/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service interface {
	saveArticleForRedactor(publish bool, newArticle models.Article) error
}

type service struct {
	repo Repository
}

func (s *service) saveArticleForRedactor(publish bool, newArticle models.Article) error {
	if newArticle.ID == uuid.Nil {
		newArticle.Reference = uuid.New()
		if publish {
			newArticle.Status = models.ArticleStatusUnderReview
			newArticle.Minor = 1 // increment minor version each time user asks for publication
		} else {
			newArticle.Status = models.ArticleStatusDraft
		}
		return s.repo.createArticle(&newArticle)
	}

	oldArticle, err := s.repo.findArticle(newArticle.ID, newArticle.RedactorID)
	if err != nil {
		return err
	}

	if oldArticle.Status != models.ArticleStatusDraft &&
		oldArticle.Status != models.ArticleStatusChangeRequested { // do NOT allow update if status is not DRAFT or CHANGE_REQUESTED
		return &fiber.Error{Code: fiber.StatusConflict, Message: fmt.Sprintf("expected one of [%s, %s], got [%s]",
			models.ArticleStatusDraft, models.ArticleStatusChangeRequested, oldArticle.Status)}
	}

	newArticle.Reference = oldArticle.Reference
	newArticle.Major = oldArticle.Major

	if publish {
		newArticle.Status = models.ArticleStatusUnderReview
		newArticle.Minor = oldArticle.Minor + 1 // increment minor version each time user asks for publication
	} else {
		newArticle.Status = oldArticle.Status
		newArticle.Minor = oldArticle.Minor
	}

	return s.repo.updateArticle(&newArticle)
}
