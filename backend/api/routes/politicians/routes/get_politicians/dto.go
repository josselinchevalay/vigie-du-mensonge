package get_politicians

import "github.com/google/uuid"

type PoliticianDTO struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
}

type ResponseDTO []PoliticianDTO
