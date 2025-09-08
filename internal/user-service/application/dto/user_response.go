package dto

import (
	"time"

	"github.com/alimahboubi/gochat/internal/user-service/domain/entities"
)

type UserResponse struct {
	ID              string     `json:"id"`
	Email           string     `json:"email"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	FullName        string     `json:"full_name"`
	IsActive        bool       `json:"is_active"`
	IsVerified      bool       `json:"is_verified"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	ActivatedAt     time.Time  `json:"activated_at,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty"`
}

func MapUserToResponse(user *entities.User) *UserResponse {
	fullName := user.FirstName() + " " + user.LastName()
	return &UserResponse{
		ID:              user.Id().Value(),
		Email:           user.Email().Value(),
		FirstName:       user.FirstName(),
		LastName:        user.LastName(),
		FullName:        fullName,
		IsActive:        user.IsActive(),
		IsVerified:      user.IsVerified(),
		CreatedAt:       user.CreatedAt(),
		UpdatedAt:       user.UpdatedAt(),
		ActivatedAt:     user.ActivatedAt(),
		EmailVerifiedAt: user.EmailVerifiedAt(),
		LastLoginAt:     user.LastLoginAt(),
	}
}
