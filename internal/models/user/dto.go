package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func NewUserDTO(u *User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		UserName:  u.UserName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type CreateUserDTO struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func (u *CreateUserDTO) Validate() error {
	if strings.TrimSpace(u.UserName) != "" {
		return errors.New("user_name is required")
	}

	if strings.TrimSpace(u.Email) != "" {
		return errors.New("email is required")
	}

	return nil
}
