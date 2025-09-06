package events

import (
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserPasswordChangedEvent struct {
	BaseDomainEvent
	userId    *valueobjects.UserId
	changedAt time.Time
}

type UserPasswordChangedEventData struct {
	UserId    string `json:"user_id"`
	ChangedAt int64  `json:"changed_at"`
}

func (e *UserPasswordChangedEvent) UserId() *valueobjects.UserId {
	return e.userId
}

func (e *UserPasswordChangedEvent) EventData() interface{} {
	return UserPasswordChangedEventData{
		UserId:    e.userId.Value(),
		ChangedAt: e.changedAt.Unix(),
	}
}

func (e *UserPasswordChangedEvent) ChangedAt() time.Time {
	return e.changedAt
}

func NewUserPasswordChangedEvent(id *valueobjects.UserId, changedAt time.Time) *UserPasswordChangedEvent {
	return &UserPasswordChangedEvent{
		BaseDomainEvent: NewBaseDomainEvent("UserPasswordChanged", id.Value(), "User", 1),
		userId:          id,
		changedAt:       changedAt,
	}
}
