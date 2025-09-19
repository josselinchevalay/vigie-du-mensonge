package sign_up

import (
	"errors"
	"vdm/core/logger"
	"vdm/core/models"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Repository interface {
	createUserAndRefreshToken(user *models.User, rft *models.RefreshToken) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) createUserAndRefreshToken(user *models.User, rft *models.RefreshToken) (err error) {
	tx := r.db.Begin()

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				logger.Error("failed to rollback transaction", logger.Err(rbErr))
			}

			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				err = gorm.ErrDuplicatedKey // email already exists
			}

			return
		}

		if cmErr := tx.Commit().Error; cmErr != nil {
			logger.Error("failed to commit transaction", logger.Err(cmErr))
			err = cmErr
		}
	}()

	if err = tx.Create(user).Error; err != nil {
		return
	}

	rft.UserID = user.ID

	if err = tx.Create(rft).Error; err != nil {
		return
	}

	return
}
