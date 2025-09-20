package locals

import (
	"vdm/core/models"

	"github.com/google/uuid"
)

type AuthedUser struct {
	ID    uuid.UUID
	Email string
}

func NewAuthedUser(u models.User) AuthedUser {
	return AuthedUser{
		ID:    u.ID,
		Email: u.Email,
	}
}
