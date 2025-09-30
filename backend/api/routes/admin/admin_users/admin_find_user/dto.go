package admin_find_user

import (
	"time"
	"vdm/core/models"
)

type ResponseDTO struct {
	Tag       string            `json:"tag"`
	CreatedAt time.Time         `json:"createdAt"`
	Roles     []models.RoleName `json:"roles,omitempty"`
}
