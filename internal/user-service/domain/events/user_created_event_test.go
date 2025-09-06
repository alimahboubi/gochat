package events_test

import (
	"testing"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestUserCreatedEvent_NewUserCreatedEvent_WithValidData_ShouldCreateEvent(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	email, _ := valueobjects.NewEmail("user@example.com")
	firstName := "John"
	lastName := "Doe"

	// Act
	userCreatedEvent := events.NewUserCreatedEvent(userId, email, firstName, lastName)

	// Assert
	assert.NotNil(t, userCreatedEvent)
	assert.Equal(t, userId, userCreatedEvent.UserId())
	assert.Equal(t, email, userCreatedEvent.Email())
	assert.Equal(t, firstName, userCreatedEvent.FirstName())
	assert.Equal(t, lastName, userCreatedEvent.LastName())

	assert.Equal(t, "UserCreated", userCreatedEvent.EventType())
	assert.Equal(t, "User", userCreatedEvent.AggregateType())
	assert.Equal(t, 1, userCreatedEvent.EventVersion())

}

func TestUserCreatedEvent_EventData_ShouldReturnStructuredData(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	email, _ := valueobjects.NewEmail("user@example.com")
	firstName := "John"
	lastName := "Doe"
	userCreatedEvent := events.NewUserCreatedEvent(userId, email, firstName, lastName)

	// Act
	eventData := userCreatedEvent.EventData()
	data, ok := eventData.(events.UserCreatedEventData)
	// Assert
	assert.NotNil(t, eventData)
	assert.True(t, ok)

	assert.Equal(t, userId.Value(), data.UserId)
	assert.Equal(t, email.Value(), data.Email)
	assert.Equal(t, firstName, data.FirstName)
	assert.Equal(t, lastName, data.LastName)
}
