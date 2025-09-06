package events_test

import (
	"testing"
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestUserEmailVerifiedEvent_NewUserEmailVerifiedEvent_WithValidData_ShouldCreateEvent(t *testing.T) {

	// Arrange
	userId := valueobjects.NewUserId()
	email, _ := valueobjects.NewEmail("test@example.com")
	verifiedAt := time.Now()

	// Act
	userEmailVerifiedEvent := events.NewUserEmailVerifiedEvent(userId, email, verifiedAt)

	// Assert
	assert.NotNil(t, userEmailVerifiedEvent)
	assert.Equal(t, email, userEmailVerifiedEvent.Email())
	assert.Equal(t, userId, userEmailVerifiedEvent.UserId())
	assert.Equal(t, verifiedAt, userEmailVerifiedEvent.VerifiedAt())

	assert.Equal(t, "UserEmailVerified", userEmailVerifiedEvent.EventType())
	assert.Equal(t, "User", userEmailVerifiedEvent.AggregateType())

}

func TestUserEmailVerifiedEvent_EventData_ShouldReturnStructuredData(t *testing.T) {
	// Arrange
	userId := valueobjects.NewUserId()
	email, _ := valueobjects.NewEmail("test@example.com")
	verifiedAt := time.Now()
	userEmailVerifiedEvent := events.NewUserEmailVerifiedEvent(userId, email, verifiedAt)

	// Act
	eventData := userEmailVerifiedEvent.EventData()
	data, ok := eventData.(events.UserEmailVerifiedEventData)

	// Assert
	assert.NotNil(t, eventData)
	assert.True(t, ok)
	assert.Equal(t, userId.Value(), data.UserId)
	assert.Equal(t, email.Value(), data.Email)
	assert.Equal(t, verifiedAt.Unix(), data.VerifiedAt)

}
