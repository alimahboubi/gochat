package events

import (
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserProfileUpdatedEvent struct {
	BaseDomainEvent
	userId       *valueobjects.UserId
	oldFirstname string
	oldLastname  string
	newFirstname string
	newLastname  string
	updatedAt    time.Time
}

type UserProfileUpdatedEventData struct {
	UserId       string `json:"userId"`
	OldFirstname string `json:"oldFirstname"`
	OldLastname  string `json:"oldLastname"`
	NewFirstname string `json:"newFirstname"`
	NewLastname  string `json:"newLastname"`
	UpdatedAt    int64  `json:"updatedAt"`
}

func NewUserProfileUpdatedEvent(id *valueobjects.UserId, oldFirstname string, oldLastname string, newFirstname string, newLastname string, updatedAt time.Time) *UserProfileUpdatedEvent {
	return &UserProfileUpdatedEvent{
		BaseDomainEvent: NewBaseDomainEvent("UserProfileUpdated", id.Value(), "User", 1),
		userId:          id,
		oldFirstname:    oldFirstname,
		oldLastname:     oldLastname,
		newFirstname:    newFirstname,
		newLastname:     newLastname,
		updatedAt:       updatedAt,
	}
}

func (e *UserProfileUpdatedEvent) OldFirstname() string {
	return e.oldFirstname
}

func (e *UserProfileUpdatedEvent) OldLastname() string {
	return e.oldLastname
}

func (e *UserProfileUpdatedEvent) NewFirstname() string {
	return e.newFirstname
}

func (e *UserProfileUpdatedEvent) NewLastname() string {
	return e.newLastname
}

func (e *UserProfileUpdatedEvent) EventData() interface{} {
	return UserProfileUpdatedEventData{
		UserId:       e.userId.Value(),
		OldFirstname: e.oldFirstname,
		OldLastname:  e.oldLastname,
		NewFirstname: e.newFirstname,
		NewLastname:  e.newLastname,
		UpdatedAt:    e.updatedAt.Unix(),
	}
}
