package valueobjects

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserId struct {
	value string
}

func NewUserId() *UserId {
	return &UserId{
		uuid.New().String(),
	}
}

func UserIdFromString(idString string) (*UserId, error) {
	err := validateUserId(idString)
	if err != nil {
		return nil, err
	}
	return &UserId{idString}, nil
}

func (userId *UserId) Value() string {
	return userId.value
}

func (userId *UserId) Equals(two *UserId) bool {
	return userId.value == two.value
}

func validateUserId(userId string) error {
	if userId == "" {
		return errors.New("invalid user id: id cannot be empty")
	}
	_, err := uuid.Parse(userId)
	if err != nil {
		return errors.Wrap(err, "invalid user id: must be a valid UUID")
	}
	return nil
}
