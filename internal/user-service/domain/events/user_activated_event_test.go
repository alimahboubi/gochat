package events_test

import (
	"testing"
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestUserActivatedEvent_NewUserActivatedEvent_EventWithValidData_ShouldCreateEvent(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	activatedAt := time.Now()

	// Act
	userActivatedEvent := events.NewUserActivatedEvent(userId, activatedAt)

	// Assert
	assert.NotNil(t, userActivatedEvent)
	assert.Equal(t, userId, userActivatedEvent.UserId())
}

func TestUserActivatedEvent_EventData_ShouldReturnStructuredEvent(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	activatedAt := time.Now()
	userActivatedEvent := events.NewUserActivatedEvent(userId, activatedAt)

	// Act
	eventData := userActivatedEvent.EventData()
	data, ok := eventData.(events.UserActivatedEventData)

	// Assert
	assert.True(t, ok)
	assert.Equal(t, userId.Value(), data.UserId)
	assert.Equal(t, activatedAt.Unix(), data.ActivatedAt)
}
