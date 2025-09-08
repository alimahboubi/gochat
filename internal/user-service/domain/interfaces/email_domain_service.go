package interfaces

import (
	"time"
)

type EmailDomainService interface {
	GenerateEmailVerificationToken() (string, error)
	IsEmailVerificationTokenValid(token string, createdAt time.Time) bool
}
