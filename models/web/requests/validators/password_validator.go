package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

// CustomPasswordValidator checks if the password contains both letters and numbers.
func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if the password contains at least one letter and one number.
	hasLetter := false
	hasNumber := false

	for _, char := range password {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", char) {
			hasLetter = true
		} else if strings.ContainsRune("0123456789", char) {
			hasNumber = true
		}

		// If both a letter and a number are found, we can stop checking.
		if hasLetter && hasNumber {
			break
		}
	}

	return hasLetter && hasNumber
}
