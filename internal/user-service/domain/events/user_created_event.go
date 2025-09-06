package events

import (
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserCreatedEvent struct {
	BaseDomainEvent
	userId    *valueobjects.UserId
	email     *valueobjects.Email
	firstName string
	lastName  string
}

type UserCreatedEventData struct {
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUserCreatedEvent(userId *valueobjects.UserId, email *valueobjects.Email, name string, lastName string) *UserCreatedEvent {
	return &UserCreatedEvent{
		BaseDomainEvent: NewBaseDomainEvent("UserCreated", userId.Value(), "User", 1),
		userId:          userId,
		email:           email,
		firstName:       name,
		lastName:        lastName,
	}
}

func (e *UserCreatedEvent) UserId() *valueobjects.UserId {
	return e.userId
}

func (e *UserCreatedEvent) Email() *valueobjects.Email {
	return e.email
}
func (e *UserCreatedEvent) FirstName() string {
	return e.firstName
}
func (e *UserCreatedEvent) LastName() string {
	return e.lastName
}

func (e *UserCreatedEvent) EventData() interface{} {
	return UserCreatedEventData{
		UserId:    e.userId.Value(),
		Email:     e.email.Value(),
		FirstName: e.firstName,
		LastName:  e.lastName,
	}
}
