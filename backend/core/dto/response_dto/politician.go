package response_dto

import (
	"vdm/core/models"

	"github.com/google/uuid"
)

type Politician struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
}

func NewPolitician(entity models.Politician) Politician {
	return Politician{
		ID:       entity.ID,
		FullName: entity.FirstName + " " + entity.LastName,
	}
}
