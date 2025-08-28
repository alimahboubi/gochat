package valueobjects

import (
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash      string
	createdAt time.Time
}

var (
	hasUppercase   = regexp.MustCompile(`[A-Z]`)
	hasLowercase   = regexp.MustCompile(`[a-z]`)
	hasNumber      = regexp.MustCompile(`[0-9]`)
	hasSpecialChar = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
)

const (
	minLength          = 8
	PasswordExpireDays = 90
)

func (p *Password) Hash() string {
	return p.hash
}

func (p *Password) IsExpired() bool {
	expireDate := p.createdAt.AddDate(0, 0, PasswordExpireDays)
	return expireDate.After(time.Now())
}

func NewPassword(password string) (*Password, error) {

	policyErr := validatePasswordPolicy(password)
	if policyErr != nil {
		return nil, policyErr
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Password{
		hash: string(hash),
	}, nil
}

func validatePasswordPolicy(password string) error {

	if password == "" {
		return errors.New("invalid password: password cannot be empty")
	}

	if len(password) < minLength {
		return errors.Errorf("invalid password: must be at least %d characters long", minLength)
	}

	if !hasUppercase.MatchString(password) {
		return errors.New("invalid password: password must contain at least one uppercase letter")
	}
	if !hasLowercase.MatchString(password) {
		return errors.New("invalid password: password must contain at least one lowercase letter")
	}
	if !hasNumber.MatchString(password) {
		return errors.New("invalid password: password must contain at least one number")
	}
	if !hasSpecialChar.MatchString(password) {
		return errors.New("invalid password: password must contain at least one special character")
	}

	return nil
}

func PasswordFromHash(hash string) (*Password, error) {
	err := validateHash(hash)
	if err != nil {
		return nil, err
	}
	return &Password{hash: hash}, nil
}

func validateHash(hash string) error {
	if hash == "" {
		return errors.New("invalid password hash: hash cannot be empty")
	}
	if !strings.HasPrefix(hash, "$2") {
		return errors.New("invalid password hash: invalid hash format")
	}

	return nil
}
