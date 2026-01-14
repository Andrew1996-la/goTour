package function

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	// Ошибки валидации имени
	errNameNotValid = errors.New("name is not valid")

	// Ошибки валидации email
	errEmailToShort     = errors.New("email is too short")
	errEmailToLong      = errors.New("email is too long")
	errEmailHasNoAtSign = errors.New("email must contain @")

	// Ошибки валидации password
	errPasswordTooShort                 = errors.New("password is too short")
	errPasswordTooLong                  = errors.New("password is too long")
	errPasswordDontContainsLatinLetters = errors.New("password must contain Latin letters")
	errPasswordDontContainsDigits       = errors.New("password must contain digits")
	errPasswordDontContainsSpecialChars = errors.New("password must contain special chars")
)

func nameIsValid(name string) error {
	if len(name) < 2 {
		return errNameNotValid
	}

	return nil
}

func emailIsValid(email string) error {
	if len(email) < 3 {
		return errEmailToShort
	}

	if len(email) > 100 {
		return errEmailToLong
	}

	if !strings.Contains(email, "@") {
		return errEmailHasNoAtSign
	}

	return nil
}

func passwordIsValid(password string) error {
	if len(password) < 8 {
		return errPasswordTooShort
	}

	if len(password) > 100 {
		return errPasswordTooLong
	}

	hasLetter := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			hasSpecial = true
		}
	}

	if !hasLetter {
		return errPasswordDontContainsLatinLetters
	}
	if !hasDigit {
		return errPasswordDontContainsDigits
	}
	if !hasSpecial {
		return errPasswordDontContainsSpecialChars
	}

	return nil
}

func validationRegistration(name, email, password string) (bool, error) {
	if errName := nameIsValid(name); errName != nil {
		return false, errName
	}
	if errEmail := emailIsValid(email); errEmail != nil {
		return false, errEmail
	}
	if errPassword := passwordIsValid(password); errPassword != nil {
		return false, errPassword
	}

	return true, nil
}

func RegistrationValidation() {
	name := "Boba"
	email := "boba@mail.ru"
	pass := "$bobaSecure1234$"

	dataIsValid, err := validationRegistration(name, email, pass)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Valid registration result: %t\n", dataIsValid)
}
