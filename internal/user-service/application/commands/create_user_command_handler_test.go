package commands_test

import (
	"context"
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/application/commands"
	"github.com/alimahboubi/gochat/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserCommandHandler_Handle_WithValidCommand_ShouldCreateUser(t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)

	email := "john@example.com"
	password := "StrongPassword1234!"
	firstName := "John"
	lastName := "Doe"

	repoMock := mocks.NewMockUserRepository(ctrl)
	userServiceMock := mocks.NewMockUserDomainService(ctrl)
	emailServiceMock := mocks.NewMockEmailDomainService(ctrl)

	userServiceMock.EXPECT().
		ValidateUserRegistration(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	emailServiceMock.EXPECT().
		GenerateEmailVerificationToken().
		Return("token-123", nil)

	repoMock.EXPECT().
		Create(gomock.Any()).
		Return(nil)

	commandHandler := commands.NewCreateUserCommandHandler(emailServiceMock, userServiceMock, repoMock)
	command, err := commands.NewCreateUserCommand(email, password, firstName, lastName)
	require.NoError(t, err)
	ctx := context.Background()

	// Act
	result, err := commandHandler.Handle(&ctx, command)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.User)
	assert.Equal(t, email, result.User.Email)
}

func TestCreateUserCommandHandler_Handle_WithExistEmail_ShouldRaiseError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := "john@example.com"
	password := "StrongPassword1234!"
	firstName := "John"
	lastName := "Doe"

	repoMock := mocks.NewMockUserRepository(ctrl)
	userServiceMock := mocks.NewMockUserDomainService(ctrl)
	emailServiceMock := mocks.NewMockEmailDomainService(ctrl)

	userServiceMock.EXPECT().
		ValidateUserRegistration(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errors.New("email is already taken"))

	commandHandler := commands.NewCreateUserCommandHandler(emailServiceMock, userServiceMock, repoMock)
	command, err := commands.NewCreateUserCommand(email, password, firstName, lastName)
	require.NoError(t, err)
	ctx := context.Background()

	// Act
	result, err := commandHandler.Handle(&ctx, command)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "email is already taken")
}
