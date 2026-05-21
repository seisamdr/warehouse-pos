package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(data interface{}) error {
	var errorMessages []string

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", err.Field()))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is not a valid email", err.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param()))
			default:
				errorMessages = append(errorMessages, err.Error())
			}
		}
		return errors.New("Validation failed: " + joinMessages(errorMessages))
	}
	return nil
}

func joinMessages(messages []string) string {
	result := ""
	for i, message := range messages {
		if i > 0 {
			result += ", "
		}
		result += message
	}
	return result
}
