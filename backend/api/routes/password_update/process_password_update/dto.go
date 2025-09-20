package process_password_update

type RequestDTO struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,password"`
}
