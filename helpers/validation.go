package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MapValidationErrors(err error) []string {
	var errorMessages []string

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			msg := mapErrorMessage(e)
			errorMessages = append(errorMessages, msg)
		}
	}
	return errorMessages
}

func mapErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}
