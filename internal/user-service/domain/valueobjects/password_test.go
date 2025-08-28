package valueobjects_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestPassword_NewPassword_WithValidPassword_ShouldCreateHashPassword(t *testing.T) {
	// Arrange
	validPassword := "SecurePassword123!"

	// Act
	password, err := valueobjects.NewPassword(validPassword)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, password)
	assert.NotEqual(t, validPassword, password.Hash())
	assert.NotEmpty(t, password.Hash())
}

func TestPassword_NewPassword_WithWeekPassword_ShouldReturnError(t *testing.T) {
	// Arrange
	weekPasswords := []struct {
		name     string
		password string
		reason   string
	}{
		{"empty password", "", "cannot be empty"},
		{"too short", "123", "at least 8 characters"},
		{"no uppercase", "password123!", "uppercase letter"},
		{"no lowercase", "PASSWORD123!", "lowercase letter"},
		{"no numbers", "Password!", "number"},
		{"no special chars", "Password123", "special character"},
	}
	for _, weekPassword := range weekPasswords {
		// Act
		pass, err := valueobjects.NewPassword(weekPassword.password)

		// Assert
		assert.Nil(t, pass)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), weekPassword.reason)
	}
}

func TestPassword_FromHash_WithValidHash_ShouldReturnPassword(t *testing.T) {
	// Arrange
	plainPassword := "SecurePassword123!"
	originalPassword, _ := valueobjects.NewPassword(plainPassword)
	hash := originalPassword.Hash()

	// Act
	password, err := valueobjects.PasswordFromHash(hash)

	// Assert
	assert.NotNil(t, password)
	assert.NoError(t, err)
	assert.Equal(t, hash, password.Hash())
}

func TestPassword_FromHash_WithInvalidHash_ShouldReturnError(t *testing.T) {
	// Arrange
	invalidHash := "invalidHash"

	// Act
	_, err := valueobjects.PasswordFromHash(invalidHash)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password hash")

}

func TestPassword_IsExpired_WithNewPassword_ShouldReturnFalse(t *testing.T) {
	// Arrange
	plainPassword := "SecurePassword123!"
	password, _ := valueobjects.NewPassword(plainPassword)

	// Act
	isExpired := password.IsExpired()

	// Assert

	assert.False(t, isExpired)
}
