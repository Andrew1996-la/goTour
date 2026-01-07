package processingErrors

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyName    = errors.New("name is empty")
	ErrAgeTooLow    = errors.New("age must be at least 18")
	ErrInvalidEmail = errors.New("invalid email")
)

func validateUser(name, email string, age int) error {
	errorList := make([]error, 0)

	if name == "" {
		errorList = append(errorList, ErrEmptyName)
	}
	if email == "" || !strings.Contains(email, "@") {
		errorList = append(errorList, ErrInvalidEmail)
	}
	if age < 18 {
		errorList = append(errorList, ErrAgeTooLow)
	}

	return errors.Join(errorList...)
}

func unitedErrors() {
	err := validateUser("в", "testmail@.com", 15)
	if errors.Is(err, ErrEmptyName) {
		fmt.Println("Name не валидно")
	}
	if errors.Is(err, ErrInvalidEmail) {
		fmt.Println("Email не валидный")
	}
	if errors.Is(err, ErrAgeTooLow) {
		fmt.Println("Age меньше 18")
	}
}
