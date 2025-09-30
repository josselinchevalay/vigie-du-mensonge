package admin_find_user

import (
	"errors"
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	findUserByTag(tag string) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) findUserByTag(tag string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("tag = ?", tag).
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Select("id", "tag", "created_at").
		First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
