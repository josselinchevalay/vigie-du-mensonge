package inquire_sign_up

type RequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}
