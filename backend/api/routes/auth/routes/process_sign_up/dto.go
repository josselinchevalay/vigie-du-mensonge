package process_sign_up

import (
	"time"
	"vdm/core/models"
)

type RequestDTO struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type ResponseDTO struct {
	AccessTokenExpiry  time.Time         `json:"accessTokenExpiry"`
	RefreshTokenExpiry time.Time         `json:"refreshTokenExpiry"`
	Roles              []models.RoleName `json:"roles,omitempty"`
}
