package locals

import (
	"time"

	"github.com/google/uuid"
)

type AccessToken struct {
	Token  string
	Expiry time.Time
}

type RefreshToken struct {
	Token  uuid.UUID
	Expiry time.Time
}
