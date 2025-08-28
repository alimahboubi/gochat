package valueobjects_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserId_NewUserId_ShouldGenerateValidId(t *testing.T) {
	//Act
	userId := valueobjects.NewUserId()

	// Assert
	assert.NotNil(t, userId, "NewUserId should not be nil")
	assert.NotEmpty(t, userId, "NewUserId should not be empty")
}

func TestUserId_FromString_WithValidUUID_ShouldCreateUserId(t *testing.T) {
	// Arrange
	validUserIdString := uuid.New().String()

	// Act
	userId, err := valueobjects.UserIdFromString(validUserIdString)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, userId)
	assert.Equal(t, validUserIdString, userId.Value())
}

func TestUserId_Equals_WithSameId_ShouldReturnTrue(t *testing.T) {
	// Arrange
	userIdString := uuid.New().String()
	userIdOne, _ := valueobjects.UserIdFromString(userIdString)
	userIdTwo, _ := valueobjects.UserIdFromString(userIdString)

	// Act
	isEqual := userIdOne.Equals(userIdTwo)

	// Assert
	assert.True(t, isEqual)
}

func TestUserId_Equals_WithDifferentId_ShouldReturnFalse(t *testing.T) {
	// Arrange
	userIdOne := valueobjects.NewUserId()
	userIdTwo := valueobjects.NewUserId()

	// Act
	isEqual := userIdOne.Equals(userIdTwo)

	// Assert
	assert.False(t, isEqual)
}
