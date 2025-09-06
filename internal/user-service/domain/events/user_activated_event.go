package events

import (
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserActivatedEvent struct {
	BaseDomainEvent
	userId      *valueobjects.UserId
	activatedAt time.Time
}

type UserActivatedEventData struct {
	UserId      string `json:"user_id"`
	ActivatedAt int64  `json:"activated_at"`
}

func NewUserActivatedEvent(userId *valueobjects.UserId, activatedAt time.Time) *UserActivatedEvent {
	return &UserActivatedEvent{
		BaseDomainEvent: NewBaseDomainEvent("UserActivated", userId.Value(), "User", 1),
		userId:          userId,
		activatedAt:     activatedAt,
	}
}

func (e *UserActivatedEvent) UserId() *valueobjects.UserId {
	return e.userId
}

func (e *UserActivatedEvent) ActivateAt() time.Time {
	return e.activatedAt
}

func (e *UserActivatedEvent) EventData() interface{} {
	return UserActivatedEventData{
		UserId:      e.userId.Value(),
		ActivatedAt: e.activatedAt.Unix(),
	}
}
