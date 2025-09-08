package services

import (
	"github.com/alimahboubi/gochat/internal/user-service/domain/interfaces"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/pkg/errors"
)

type UserDomainService struct {
	repo interfaces.UserRepository
}

func NewUserDomainService(repo interfaces.UserRepository) *UserDomainService {
	return &UserDomainService{repo}
}

func (s *UserDomainService) IsEmailUnique(email *valueobjects.Email) (bool, error) {
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return false, errors.Wrap(err, "failed to check if email is unique")
	}

	return !exists, nil
}

func (s *UserDomainService) ValidateUserRegistration(email *valueobjects.Email, password *valueobjects.Password, firstName, lastName string) error {
	exists, err := s.IsEmailUnique(email)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("email is already taken")
	}
	return nil
}
