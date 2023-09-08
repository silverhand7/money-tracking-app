package validators

import "github.com/go-playground/validator"

type ValidType struct {
	Type string `validate:"in=i e"` // income, expense
}

func ValidateType(fl validator.FieldLevel) bool {
	validTypes := []string{"i", "e"}
	value := fl.Field().String()

	for _, validType := range validTypes {
		if value == validType {
			return true
		}
	}

	return false
}
