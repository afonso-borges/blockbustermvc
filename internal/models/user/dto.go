package models

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserDTO(u *User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		UserName:  u.UserName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}

type CreateUserDTO struct {
	UserName string `json:"user_name" binding:"required,min=4,max=100"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserDTO struct {
	UserName string `json:"user_name" binding:"required,min=4,max=100"`
	Email    string `json:"email" binding:"required,email"`
}
