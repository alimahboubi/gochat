package commands

import (
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/pkg/errors"
)

type CreateUserCommand struct {
	Email     *valueobjects.Email
	Password  *valueobjects.Password
	Firstname string
	Lastname  string
}

func NewCreateUserCommand(email, password, firstname, lastname string) (*CreateUserCommand, error) {

	emailObj, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordObj, err := valueobjects.NewPassword(password)
	if err != nil {
		return nil, err
	}

	if firstname == "" {
		return nil, errors.New("firstname cannot be empty")
	}

	if lastname == "" {
		return nil, errors.New("lastname cannot be empty")
	}

	return &CreateUserCommand{
		Email:     emailObj,
		Password:  passwordObj,
		Firstname: firstname,
		Lastname:  lastname,
	}, nil
}

func (c *CreateUserCommand) CommandType() string {
	return "CreateUser"
}
