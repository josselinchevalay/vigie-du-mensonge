package locals

import (
	"slices"
	"vdm/core/models"

	"github.com/google/uuid"
)

type AuthedUser struct {
	ID            uuid.UUID
	Email         string
	EmailVerified bool
	Roles         []models.RoleName
}

func (a AuthedUser) HasRole(role models.RoleName) bool {
	return slices.Contains(a.Roles, role)
}
