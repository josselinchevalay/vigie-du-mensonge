package process_password_update

import "github.com/google/uuid"

type RequestDTO struct {
	Token       uuid.UUID `json:"token" validate:"required"`
	NewPassword string    `json:"newPassword" validate:"required,password"`
}
