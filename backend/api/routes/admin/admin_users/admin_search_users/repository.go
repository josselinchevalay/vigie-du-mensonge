package admin_search_users

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	searchUsersByTag(query string) ([]models.User, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) searchUsersByTag(query string) ([]models.User, error) {
	var users []models.User

	if err := r.db.Where("tag ILIKE ?", "%"+query+"%").
		Limit(20).
		Select("tag").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
