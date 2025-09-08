package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id                 string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email              string     `gorm:"type:varchar(255);uniqueIndex;not null;unique" json:"email"`
	PasswordHash       string     `gorm:"type:varchar(255);not null" json:"password_hash"`
	FirstName          string     `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(50);not null" json:"last_name"`
	IsActive           bool       `gorm:"not null;default:false" json:"is_active"`
	IsVerified         bool       `gorm:"not null;default:false" json:"is_verified"`
	CreatedAt          time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"null" json:"updated_at"`
	DeactivatedAt      *time.Time `gorm:"null"  json:"deactivated_at"`
	EmailVerifiedAt    *time.Time `gorm:"null"  json:"email_verified_at"`
	LastLoginAt        *time.Time `gorm:"null"  json:"last_login_at"`
	DeactivationReason string     `gorm:"type:text" json:"deactivation_reason"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = &now
	}
	return nil
}
