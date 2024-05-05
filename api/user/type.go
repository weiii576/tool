package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_EMAIL_ALREADY_USED        = "email already used"
	MESSAGE_FAILED_CREATE_USER        = "failed create user"

	// Success
	MESSAGE_SUCCESS_CREATE_USER = "success create user"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Birthday string `form:"birthday"`
}
