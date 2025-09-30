package moderator_save_review

import (
	"fmt"
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	createReviewAndUpdateArticle(review *models.ArticleReview) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) createReviewAndUpdateArticle(review *models.ArticleReview) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var article models.Article

		if err := tx.Where("id = ? AND moderator_id = ? AND status = ?", review.ArticleID, review.ModeratorID, models.ArticleStatusUnderReview).
			Select("id", "major").
			First(&article).Error; err != nil {
			return fmt.Errorf("failed to find Article{ID=%s ModeratorID=%s Status=%s}: %v",
				review.ArticleID, review.ModeratorID, models.ArticleStatusUnderReview, err)
		}

		if err := tx.Create(review).Error; err != nil {
			return fmt.Errorf("failed to create new review: %v", err)
		}

		updates := map[string]any{"status": review.Decision}

		if review.Decision == models.ArticleStatusPublished {
			updates["major"] = article.Major + 1 // increment major version each time an article is published
			updates["minor"] = 0
		}

		if err := tx.Model(&models.Article{}).
			Where("id = ?", article.ID).
			Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to updates Article{ID=%s ModeratorID=%s}: %v",
				review.ArticleID, review.ModeratorID, err)
		}

		return nil
	})
}
