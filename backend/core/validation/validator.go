package validation

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var customValidator = initValidator()

type CustomValidator struct {
	validate *validator.Validate
}

func initValidator() *CustomValidator {
	v := validator.New()
	if err := v.RegisterValidation("password", validatePassword); err != nil {
		log.Fatal(err)
	}
	if err := v.RegisterValidation("email", validateEmail); err != nil {
		log.Fatal(err)
	}
	if err := v.RegisterValidation("username", validateUsername); err != nil {
		log.Fatal(err)
	}
	return &CustomValidator{validate: v}
}

func Validate(i any) error {
	if err := customValidator.validate.Struct(i); err != nil {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}
	return nil
}
