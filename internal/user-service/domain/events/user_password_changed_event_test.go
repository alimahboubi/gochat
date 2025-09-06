package events_test

import (
	"testing"
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestUserPasswordChangedEvent_NewPasswordChangedEvent_WithValidData_ShouldCreate(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	passwordChangedAt := time.Now()

	// Act
	passwordChangedEvent := events.NewUserPasswordChangedEvent(userId, passwordChangedAt)

	// Assert
	assert.NotNil(t, passwordChangedEvent)
	assert.Equal(t, userId, passwordChangedEvent.UserId())
	assert.Equal(t, passwordChangedAt, passwordChangedEvent.ChangedAt())
	assert.Equal(t, "UserPasswordChanged", passwordChangedEvent.EventType())
	assert.Equal(t, "User", passwordChangedEvent.AggregateType())

}

func TestUserPasswordChangedEvent_EventData_ShouldCreate(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	passwordChangedAt := time.Now()
	passwordChangedEvent := events.NewUserPasswordChangedEvent(userId, passwordChangedAt)

	// Act
	userEvent := passwordChangedEvent.EventData()
	eventData, ok := userEvent.(events.UserPasswordChangedEventData)

	// Assert
	assert.True(t, ok)
	assert.NotNil(t, eventData)
	assert.Equal(t, userId.Value(), eventData.UserId)
	assert.Equal(t, passwordChangedAt.Unix(), eventData.ChangedAt)
}
