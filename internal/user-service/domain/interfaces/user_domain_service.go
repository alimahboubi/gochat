package interfaces

import "github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"

type UserDomainService interface {
	IsEmailUnique(email *valueobjects.Email) (bool, error)
	ValidateUserRegistration(email *valueobjects.Email, password *valueobjects.Password, firstName, lastName string) error
	//GetUserDomainStatistics() (*services.UserDomainStatistics, error)
}
