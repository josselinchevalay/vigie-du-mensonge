package get_politicians

import (
	"vdm/core/models"

	"gorm.io/gorm"
)

type Repository interface {
	getPoliticians() ([]models.Politician, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getPoliticians() ([]models.Politician, error) {
	var politicians []models.Politician

	if err := r.db.Model(&models.Politician{}).
		Select("id, first_name, last_name").
		Find(&politicians).Error; err != nil {
		return nil, err
	}

	return politicians, nil
}
