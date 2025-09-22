package process_email_verification

import "github.com/google/uuid"

type RequestDTO struct {
	Token uuid.UUID `json:"token" validate:"required"`
}
