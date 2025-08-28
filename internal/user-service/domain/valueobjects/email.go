package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(email string) (*Email, error) {
	err := validateEmail(email)
	if err != nil {
		return nil, err
	}
	return &Email{strings.ToLower(email)}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func validateEmail(email string) error {
	if email == "" {
		return errors.New("invalid email: email cannot be empty")
	}

	if !validateSpaceInEmail(email) {
		return errors.New("invalid email: email cannot contain white spaces")
	}

	if !emailRegex.MatchString(email) {
		return errors.New("invalid email:  format is incorrect")
	}
	return nil
}

func validateSpaceInEmail(email string) bool {
	trimEmail := strings.TrimSpace(email)
	return email == trimEmail
}
