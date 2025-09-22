package locals

import (
	"vdm/core/models"

	"github.com/google/uuid"
)

type AuthedUser struct {
	ID            uuid.UUID
	Email         string
	EmailVerified bool
	Roles         []models.RoleName
}
