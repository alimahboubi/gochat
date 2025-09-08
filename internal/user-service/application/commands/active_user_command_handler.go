package commands

import (
	"context"

	"github.com/alimahboubi/gochat/internal/user-service/application/dto"
	"github.com/alimahboubi/gochat/internal/user-service/domain/interfaces"
	"github.com/pkg/errors"
)

type ActiveUserCommandHandler struct {
	userRepo interfaces.UserRepository
}

func NewActiveUserCommandHandler(userRepo interfaces.UserRepository) *ActiveUserCommandHandler {
	return &ActiveUserCommandHandler{
		userRepo: userRepo,
	}
}

func (c *ActiveUserCommandHandler) Handle(ctx *context.Context, command *ActiveUserCommand) (*dto.ActiveUserResponse, error) {
	user, err := c.userRepo.FindById(command.UserId)
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := user.Active(); err != nil {
		return nil, err
	}

	if err := c.userRepo.Update(user); err != nil {
		return nil, err
	}
	return &dto.ActiveUserResponse{
		ActivatedAt: user.ActivatedAt(),
	}, nil
}
