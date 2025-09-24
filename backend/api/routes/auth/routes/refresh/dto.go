package refresh

import (
	"time"
	"vdm/core/models"
)

type ResponseDTO struct {
	AccessTokenExpiry  time.Time         `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time         `json:"refreshTokenExpiry"`
	Roles              []models.RoleName `json:"roles,omitempty"`
}
