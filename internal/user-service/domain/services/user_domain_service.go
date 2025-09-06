package services

import (
	"github.com/alimahboubi/gochat/internal/user-service/domain/repositories"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/pkg/errors"
)

type UserDomainService struct {
	repo repositories.UserRepository
}

func NewUserDomainService(repo repositories.UserRepository) *UserDomainService {
	return &UserDomainService{repo}
}

func (s *UserDomainService) IsEmailUnique(email *valueobjects.Email) (bool, error) {
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return false, errors.Wrap(err, "failed to check if email is unique")
	}

	return !exists, nil
}
