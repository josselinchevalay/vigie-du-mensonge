package inquire_password_update

type RequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}
