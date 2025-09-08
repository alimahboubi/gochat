package commands

import (
	"context"

	"github.com/alimahboubi/gochat/internal/user-service/application/dto"
	"github.com/alimahboubi/gochat/internal/user-service/application/interfaces"
	"github.com/alimahboubi/gochat/internal/user-service/domain/entities"
	domainInterfaces "github.com/alimahboubi/gochat/internal/user-service/domain/interfaces"
	"github.com/pkg/errors"
)

type CreateUserCommandHandler struct {
	emailService domainInterfaces.EmailDomainService
	userService  domainInterfaces.UserDomainService
	userRepo     domainInterfaces.UserRepository
}

func NewCreateUserCommandHandler(emailService domainInterfaces.EmailDomainService, userService domainInterfaces.UserDomainService, userRepo domainInterfaces.UserRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		emailService: emailService,
		userService:  userService,
		userRepo:     userRepo,
	}
}

func (c *CreateUserCommandHandler) Handle(ctx *context.Context, command *CreateUserCommand) (*dto.CreateUserResponse, error) {
	err := c.userService.ValidateUserRegistration(command.Email, command.Password, command.Firstname, command.Lastname)
	if err != nil {
		return nil, err
	}

	user, err := entities.NewUser(command.Email, command.Password, command.Firstname, command.Lastname)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user entity")
	}

	verificationToken, err := c.emailService.GenerateEmailVerificationToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate verification token")
	}

	if err := c.userRepo.Create(user); err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &dto.CreateUserResponse{
		User:                 dto.MapUserToResponse(user),
		VerificationToken:    verificationToken,
		VerificationRequired: true,
	}, nil

}

var _ interfaces.CommandHandler[*CreateUserCommand, *dto.CreateUserResponse] = (*CreateUserCommandHandler)(nil)
