package entities_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/entities"
	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_NewUser_WithValidData_ShouldCreateUser(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	// Act
	user, err := entities.NewUser(email, password, name, lastName)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.Id())
	assert.Equal(t, email, user.Email())
	assert.Equal(t, name, user.FirstName())
	assert.Equal(t, lastName, user.LastName())
	assert.False(t, user.IsActive())
	assert.False(t, user.IsVerified())
}

func TestUser_NewUser_WithValidData_ShouldRaisedUserCreatedEvent(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, password, name, lastName)

	// Act
	userEvents := user.DomainEvents()

	// Assert
	assert.NotNil(t, userEvents)
	assert.Len(t, userEvents, 1)
	userCreatedEvent, ok := userEvents[0].(*events.UserCreatedEvent)
	assert.True(t, ok)
	assert.NotNil(t, userCreatedEvent)
	assert.Equal(t, email, userCreatedEvent.Email())
	assert.Equal(t, name, userCreatedEvent.FirstName())
}

func TestUser_NewUser_WithInvalidData_ShouldReturnError(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	testCases := []struct {
		name      string
		email     *valueobjects.Email
		password  *valueobjects.Password
		firstname string
		lastname  string
	}{
		{"Invalid email", nil, password, "John", "Dou"},
		{"Invalid password", email, nil, "John", "Dou"},
		{"empty firstname", email, password, "", "Dou"},
		{"empty lastname", email, password, "john", ""},
		{"whitespace first name", email, password, "   ", "Doe"},
		{"whitespace last name", email, password, "John", "   "},
	}

	// Act & Assert
	for _, tc := range testCases {
		// Act
		user, err := entities.NewUser(tc.email, tc.password, tc.firstname, tc.lastname)

		// Assert
		assert.Nil(t, user)
		require.Error(t, err)
	}
}

func TestUser_Active_WithInactive_ShouldActiveUser(t *testing.T) {
	// Arrange
	user := createValidUser(t)
	_ = user.VerifyEmail()

	// Act
	err := user.Active()

	// Assert
	assert.NoError(t, err)
	assert.True(t, user.IsActive())

}

func TestUser_Active_WithInactive_ShouldRaisedUserActivatedEvent(t *testing.T) {
	// Arrange
	user := createValidUser(t)
	_ = user.VerifyEmail()

	// Act
	_ = user.Active()
	userEvents := user.DomainEvents()

	// Assert
	assert.Len(t, userEvents, 3)
	userActivatedEvent, ok := userEvents[2].(*events.UserActivatedEvent)
	assert.True(t, ok)
	assert.NotNil(t, userActivatedEvent)

}

func TestUser_Active_WithActive_ShouldReturnError(t *testing.T) {
	// Arrange
	user := createValidUser(t)
	_ = user.Active()

	// Act
	err := user.Active()

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user is already active")

}

func TestUser_Activate_WithUnverifiedEmail_ShouldReturnError(t *testing.T) {
	// Arrange
	user := createValidUser(t)

	// Act
	err := user.Active()

	// Assert
	require.Error(t, err)
	assert.False(t, user.IsActive())
	assert.Contains(t, err.Error(), "email must be verified before activation")

}

func TestUser_VerifyEmail_WithVerified_ShouldReturnError(t *testing.T) {
	// Arrange
	user := createValidUser(t)
	_ = user.VerifyEmail()

	// Act
	err := user.VerifyEmail()

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user is already verified")
}

func TestUser_VerifyEmail_WithUnverified_ShouldVerifiedEmail(t *testing.T) {
	// Arrange
	user := createValidUser(t)

	//Act
	err := user.VerifyEmail()

	// Assert
	assert.NoError(t, err)
	assert.True(t, user.IsVerified())
}

func TestUser_VerifyEmail_WithUnverified_ShouldRaisedUserEmailVerifiedEvent(t *testing.T) {
	// Arrange
	user := createValidUser(t)

	//Act
	_ = user.VerifyEmail()
	userEvents := user.DomainEvents()

	// Assert
	event, ok := userEvents[1].(*events.UserEmailVerifiedEvent)
	assert.True(t, ok)
	assert.NotNil(t, event)
}

func TestUser_ChangePassword_WithValidPassword_ShouldPasswordChanged(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	oldPassword, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, oldPassword, name, lastName)
	newPassword, _ := valueobjects.NewPassword("NewSecurePassword123!")

	// Act
	err := user.ChangePassword(newPassword)

	// Assert
	assert.Nil(t, err)
	assert.True(t, user.Password().VerifiedPassword("NewSecurePassword123!"))
	assert.False(t, user.Password().VerifiedPassword("StrongPassword123!"))
}

func TestUser_ChangePassword_WithNilPassword_ShouldReturnError(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	oldPassword, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, oldPassword, name, lastName)

	// Act
	err := user.ChangePassword(nil)

	// Assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "password cannot be nil")
}

func TestUser_ChangePassword_WithValidPassword_ShouldRaisedUserPasswordChangedEvent(t *testing.T) {
	// Arrange
	user := createValidUser(t)
	newPassword, _ := valueobjects.NewPassword("NewSecurePassword123!")

	//Act
	_ = user.ChangePassword(newPassword)
	userEvents := user.DomainEvents()

	// Assert
	event, ok := userEvents[1].(*events.UserPasswordChangedEvent)
	assert.True(t, ok)
	assert.NotNil(t, event)
}

func TestUser_UpdateProfile_WithValidData_ShouldUpdated(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, password, name, lastName)
	newFirstName := "Jon"
	newLastName := "Peterson"

	// Act
	err := user.UpdateProfile(newFirstName, newLastName)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, newFirstName, user.FirstName())
	assert.Equal(t, newLastName, user.LastName())
}

func TestUser_UpdateProfile_WithInvalidFirstname_ShouldReturnError(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, password, name, lastName)
	newFirstName := ""
	newLastName := "Peterson"

	// Act
	err := user.UpdateProfile(newFirstName, newLastName)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "first name is invalid")
}

func TestUser_UpdateProfile_WithInvalidLastName_ShouldReturnError(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, password, name, lastName)
	newFirstName := "Jon"
	newLastName := ""

	// Act
	err := user.UpdateProfile(newFirstName, newLastName)

	// Assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "last name is invalid")
}

func TestUser_UpdateProfile_WithValidData_ShouldRaisedUserProfileUpdatedEvent(t *testing.T) {
	// Arrange
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, _ := entities.NewUser(email, password, name, lastName)
	newFirstName := "Jon"
	newLastName := "Peterson"
	_ = user.UpdateProfile(newFirstName, newLastName)

	// Act
	userEvents := user.DomainEvents()

	// Assert
	event, ok := userEvents[1].(*events.UserProfileUpdatedEvent)
	assert.True(t, ok)
	assert.NotNil(t, event)
}

func createValidUser(t *testing.T) *entities.User {
	email, _ := valueobjects.NewEmail("john@gmail.com")
	password, _ := valueobjects.NewPassword("StrongPassword123!")
	name := "John"
	lastName := "Doe"
	user, err := entities.NewUser(email, password, name, lastName)
	assert.NoError(t, err)
	return user
}
