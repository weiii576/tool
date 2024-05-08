package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/weiii576/tool/internal"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_EMAIL_ALREADY_USED        = "email already used"
	MESSAGE_FAILED_CREATE_USER        = "failed create user"
	MESSAGE_FAILED_GET_USERS          = "failed get users"

	// Success
	MESSAGE_SUCCESS_CREATE_USER = "success create user"
	MESSAGE_SUCCESS_GET_USERS   = "success get users"
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

type Profile struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
}

type GetUsersResponse struct {
	Users      []Profile                   `json:"users"`
	Pagination internal.PaginationResponse `json:"pagination"`
}
