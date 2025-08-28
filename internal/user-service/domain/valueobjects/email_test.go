package valueobjects_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmail_NewEmail_WithValidEmail_ShouldCreateEmail(t *testing.T) {
	//Arrange
	validEmail := "test@gmail.com"
	//Act
	email, err := valueobjects.NewEmail(validEmail)
	//Assert
	require.NoError(t, err)
	assert.Equal(t, validEmail, email.Value())
}

func TestEmail_NewEmail_WithInvalidEmail_ShouldReturnError(t *testing.T) {
	//Arrange
	invalidEmails := []struct {
		name  string
		email string
	}{
		{"empty", ""},
		{"missing @", "testgmail.com"},
		{"missing domain", "test@"},
		{"missing dot", "test@gmail"},
		{"missing local part", "@gmail.com"},
		{"contain space", "test @gmail.com"},
	}

	for _, invalidEmail := range invalidEmails {
		//Act
		email, err := valueobjects.NewEmail(invalidEmail.email)
		//Assert
		require.Error(t, err)
		assert.Nil(t, email)
		assert.Contains(t, err.Error(), "invalid email")
	}
}

func TestEmail_Equals_SameEmails_ShouldReturnTrue(t *testing.T) {
	//Arrange
	emailOne, _ := valueobjects.NewEmail("testEmail@gmail.com")
	emailTwo, _ := valueobjects.NewEmail("testEmail@gmail.com")

	//Act
	result := emailOne.Equals(emailTwo)

	//Assert
	assert.True(t, result)
}

func TestEmail_Equals_DifferentEmails_ShouldReturnFalse(t *testing.T) {
	//Arrange
	emailOne, _ := valueobjects.NewEmail("testEmail@gmail.com")
	emailTwo, _ := valueobjects.NewEmail("othertestEmail@gmail.com")

	//Act
	result := emailOne.Equals(emailTwo)

	//Assert
	assert.False(t, result)
}
