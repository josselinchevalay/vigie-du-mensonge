package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
