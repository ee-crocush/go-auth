package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func ValidateRequest(req interface{}) error {
	if err := validate.Struct(req); err != nil {
		return errors.New(parseValidationError(err))
	}
	return nil
}

func parseValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("Field %s is required", e.Field()))
			case "email":
				messages = append(messages, fmt.Sprintf("Field %s must be a valid email", e.Field()))
			case "min":
				messages = append(messages, fmt.Sprintf("Field %s must have at least %s characters", e.Field(), e.Param()))
			case "max":
				messages = append(messages, fmt.Sprintf("Field %s must have no more than %s characters", e.Field(), e.Param()))
			default:
				messages = append(messages, fmt.Sprintf("Field %s is invalid", e.Field()))
			}
		}
		return strings.Join(messages, "; ")
	}
	return "Validation error not found"
}
