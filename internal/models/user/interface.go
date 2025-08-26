package models

import "github.com/google/uuid"

type IUserService interface {
	CreateUser(user *CreateUserDTO) error
	GetUser(id uuid.UUID) (*UserDTO, error)
	GetAllUsers() ([]*UserDTO, error)
	UpdateUser(id uuid.UUID, user *UserDTO) error
	DeleteUser(id uuid.UUID) error
}

type IUserRepository interface {
	CreateUser(user *CreateUserDTO) error
	GetUser(id uuid.UUID) (*UserDTO, error)
	GetAllUsers() ([]*UserDTO, error)
	UpdateUser(id uuid.UUID, user *UserDTO) error
	DeleteUser(id uuid.UUID) error
}
