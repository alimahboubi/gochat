package services

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/pkg/errors"
)

type EmailDomainService struct{}

func NewEmailDomainService() *EmailDomainService {
	return &EmailDomainService{}
}

func (es *EmailDomainService) GenerateEmailVerificationToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", errors.Wrap(err, "failed to generate verification token")
	}
	return hex.EncodeToString(bytes), nil
}

func (es *EmailDomainService) IsEmailVerificationTokenValid(token string, createdAt time.Time) bool {
	expireTime := createdAt.Add(24 * time.Hour)
	return time.Now().Before(expireTime) && len(token) == 64
}
