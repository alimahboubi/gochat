package events

import (
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserEmailVerifiedEvent struct {
	BaseDomainEvent
	userId     *valueobjects.UserId
	email      *valueobjects.Email
	verifiedAt time.Time
}

type UserEmailVerifiedEventData struct {
	UserId     string `json:"userId"`
	Email      string `json:"email"`
	VerifiedAt int64  `json:"verifiedAt"`
}

func (e *UserEmailVerifiedEvent) Email() *valueobjects.Email {
	return e.email
}

func (e *UserEmailVerifiedEvent) UserId() *valueobjects.UserId {
	return e.userId
}
func (e *UserEmailVerifiedEvent) VerifiedAt() time.Time {
	return e.verifiedAt
}

func (e *UserEmailVerifiedEvent) EventData() interface{} {
	return UserEmailVerifiedEventData{
		UserId:     e.userId.Value(),
		Email:      e.email.Value(),
		VerifiedAt: e.verifiedAt.Unix(),
	}
}

func NewUserEmailVerifiedEvent(id *valueobjects.UserId, email *valueobjects.Email, verifiedAt time.Time) *UserEmailVerifiedEvent {
	return &UserEmailVerifiedEvent{
		BaseDomainEvent: NewBaseDomainEvent("UserEmailVerified", id.Value(), "User", 1),
		userId:          id,
		email:           email,
		verifiedAt:      verifiedAt,
	}
}
