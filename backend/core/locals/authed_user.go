package locals

import "github.com/google/uuid"

type AuthedUser struct {
	ID    uuid.UUID
	Roles []string
}
