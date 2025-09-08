package commands

import "github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"

type ActiveUserCommand struct {
	UserId *valueobjects.UserId
}

func (c *ActiveUserCommand) CommandType() string {
	return "ActiveUserCommand"
}
