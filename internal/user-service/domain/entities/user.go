package entities

import (
	"strings"
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/events"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
	"github.com/pkg/errors"
)

type User struct {
	id         *valueobjects.UserId
	email      *valueobjects.Email
	password   *valueobjects.Password
	firstName  string
	lastName   string
	isActive   bool
	isVerified bool
	events     []events.DomainEvent

	createdAt         time.Time
	activateAt        time.Time
	emailVerifiedAt   time.Time
	passwordChangedAt time.Time
	profileChangedAt  time.Time
	lastLoginAt       time.Time
}

func NewUser(email *valueobjects.Email, password *valueobjects.Password, firstName, lastname string) (*User, error) {
	err := validateValues(email, password, firstName, lastname)
	if err != nil {
		return nil, err
	}
	user := &User{
		id:         valueobjects.NewUserId(),
		email:      email,
		password:   password,
		firstName:  firstName,
		lastName:   lastname,
		isActive:   false,
		isVerified: false,
		createdAt:  time.Now(),
	}

	userCreatedEvent := events.NewUserCreatedEvent(user.Id(), email, firstName, lastname)
	user.events = []events.DomainEvent{userCreatedEvent}

	return user, nil
}

func validateValues(email *valueobjects.Email, password *valueobjects.Password, name string, lastname string) error {
	if email == nil {
		return errors.New("email is nil")
	}
	if password == nil {
		return errors.New("password is nil")
	}
	if name == "" || strings.TrimSpace(name) != name {
		return errors.New("name is invalid")
	}
	if lastname == "" || strings.TrimSpace(lastname) != lastname {
		return errors.New("lastname is invalid")
	}
	return nil
}

func (u *User) Id() *valueobjects.UserId {
	return u.id
}

func (u *User) Email() *valueobjects.Email {
	return u.email
}
func (u *User) Password() *valueobjects.Password {
	return u.password
}
func (u *User) FirstName() string {
	return u.firstName
}
func (u *User) LastName() string {
	return u.lastName
}

func (u *User) IsActive() bool {
	return u.isActive
}

func (u *User) IsVerified() bool {
	return u.isVerified
}

func (u *User) DomainEvents() []events.DomainEvent {
	return u.events
}

func (u *User) addDomainEvent(e events.DomainEvent) {
	u.events = append(u.events, e)
}

func (u *User) ClearDomainEvents() {
	u.events = make([]events.DomainEvent, 0)
}

func (u *User) Active() error {
	err := u.validateCanBeActivated()
	if err != nil {
		return err
	}
	now := time.Now()
	u.isActive = true
	u.activateAt = now
	userActivatedEvent := events.NewUserActivatedEvent(u.id, now)
	u.addDomainEvent(userActivatedEvent)
	return nil
}

func (u *User) validateCanBeActivated() error {
	if u.isActive {
		return errors.New("user is already active")
	}
	if !u.isVerified {
		return errors.New("email must be verified before activation")
	}
	return nil
}

func (u *User) VerifyEmail() error {
	if u.isVerified {
		return errors.New("user is already verified")
	}
	now := time.Now()
	u.isVerified = true
	u.emailVerifiedAt = now
	emailActivatedEvent := events.NewUserEmailVerifiedEvent(u.id, u.email, now)
	u.addDomainEvent(emailActivatedEvent)
	return nil
}

func (u *User) ChangePassword(password *valueobjects.Password) error {
	if password == nil {
		return errors.New("password cannot be nil")
	}
	now := time.Now()
	u.password = password
	u.passwordChangedAt = now
	passwordChangedEvent := events.NewUserPasswordChangedEvent(u.id, now)
	u.addDomainEvent(passwordChangedEvent)
	return nil
}

func (u *User) UpdateProfile(firstName string, lastName string) error {
	err := u.validateProfileUpdate(firstName, lastName)
	if err != nil {
		return err
	}
	now := time.Now()
	oldFirstName := u.firstName
	oldLastName := u.lastName
	u.firstName = firstName
	u.lastName = lastName
	u.profileChangedAt = now

	profileUpdatedEvent := events.NewUserProfileUpdatedEvent(u.id, oldFirstName, oldLastName, firstName, lastName, now)
	u.addDomainEvent(profileUpdatedEvent)

	return nil
}

func (u *User) validateProfileUpdate(firstname, lastname string) error {
	if firstname == "" {
		return errors.New("first name is invalid")
	}
	if lastname == "" {
		return errors.New("last name is invalid")
	}
	return nil
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() *time.Time {
	return &u.profileChangedAt
}

func (u *User) ActivatedAt() time.Time {
	return u.activateAt
}

func (u *User) EmailVerifiedAt() *time.Time {
	return &u.emailVerifiedAt
}

func (u *User) LastLoginAt() *time.Time {
	return &u.lastLoginAt
}
