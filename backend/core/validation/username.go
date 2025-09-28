package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const minUsernameLen = 2
const maxUsernameLen = 20

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	if len(username) < minUsernameLen || len(username) > maxUsernameLen {
		return false
	}

	usernameRegex := regexp.MustCompile(`^[a-z]+$`)

	return usernameRegex.MatchString(username)
}
