package events_test

import (
	"testing"
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestUserProfileUpdatedEvent_NewUserProfileUpdatedEvent_WithValidData_ShouldCreated(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	oldFirstName := "John"
	oldLastName := "Doe"
	newFirstName := "Jane"
	newLastName := "Peterson"
	updatedAt := time.Now()

	// Act
	userProfileUpdatedEvent := events.NewUserProfileUpdatedEvent(userId, oldFirstName, oldLastName, newFirstName, newLastName, updatedAt)

	// Assert
	assert.NotNil(t, userProfileUpdatedEvent)
	assert.Equal(t, oldFirstName, userProfileUpdatedEvent.OldFirstname())
	assert.Equal(t, oldLastName, userProfileUpdatedEvent.OldLastname())
	assert.Equal(t, newFirstName, userProfileUpdatedEvent.NewFirstname())
	assert.Equal(t, newLastName, userProfileUpdatedEvent.NewLastname())
}

func TestUserProfileUpdatedEvent_EventData_ShouldReturnStructuredData(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	oldFirstName := "John"
	oldLastName := "Doe"
	newFirstName := "Jane"
	newLastName := "Peterson"
	updatedAt := time.Now()
	userProfileUpdatedEvent := events.NewUserProfileUpdatedEvent(userId, oldFirstName, oldLastName, newFirstName, newLastName, updatedAt)

	// Act
	userEvent := userProfileUpdatedEvent.EventData()
	eventData, ok := userEvent.(events.UserProfileUpdatedEventData)

	// Assert
	assert.True(t, ok)
	assert.NotNil(t, eventData)
	assert.Equal(t, userId.Value(), eventData.UserId)
	assert.Equal(t, oldFirstName, eventData.OldFirstname)
	assert.Equal(t, oldLastName, eventData.OldLastname)
	assert.Equal(t, newFirstName, eventData.NewFirstname)
	assert.Equal(t, newLastName, eventData.NewLastname)
	assert.Equal(t, updatedAt.Unix(), eventData.UpdatedAt)

}
