package validators

import "github.com/go-playground/validator"

type ValidType struct {
	Type string `validate:"in=I E"` // income, expense
}

func ValidateType(fl validator.FieldLevel) bool {
	validTypes := []string{"I", "E"}
	value := fl.Field().String()

	for _, validType := range validTypes {
		if value == validType {
			return true
		}
	}

	return false
}
